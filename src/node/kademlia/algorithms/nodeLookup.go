package algorithms

import (
	"D7024E/environment"
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/kademlia"
	"D7024E/node/kademlia/kademliaSort"

	"sort"
	"sync"
)

type PingRpc func(contact.Contact, contact.Contact) bool
type FindNodeRPC func(contact.Contact, contact.Contact, id.KademliaID) ([]contact.Contact, error)

// Node lookup initiator.
func NodeLookup(targetID id.KademliaID) []contact.Contact {
	batch := kademlia.GetInstance().RoutingTable.FindClosestContacts(&targetID, environment.Alpha)
	return NodeLookupRec(targetID, batch, rpc.FindNodeRequest, rpc.Ping)
}

// Algorithm for Node lookup.
func NodeLookupRec(targetID id.KademliaID, batch []contact.Contact, findNode FindNodeRPC, ping PingRpc) []contact.Contact {
	batch = getAllDistances(batch)
	newBatch := findNodes(targetID, batch, findNode)
	updatedBatch := mergeBatch(newBatch)
	updatedBatch = removeDuplicates(updatedBatch)
	updatedBatch = removeDeadNodes(updatedBatch, ping)
	updatedBatch = getAllDistances(updatedBatch)
	updatedBatch = kademliaSort.SortContacts(updatedBatch)
	updatedBatch = resize(updatedBatch)
	if isSame(batch, updatedBatch) {
		return updatedBatch
	} else {
		return NodeLookupRec(targetID, updatedBatch, findNode, ping)
	}
}

// Updates the distances in "batch" to be the distances to the current node
// then returns the new batch.
func getAllDistances(batch []contact.Contact) []contact.Contact {
	for i := 0; i < len(batch); i++ {
		relativeDistance := *batch[i].ID.CalcDistance(contact.GetInstance().ID)
		batch[i].SetDistance(&relativeDistance)
	}
	return batch
}

// min value of a and b.
func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

// Find all nodes from the know contacts in batch.
func findNodes(targetID id.KademliaID, batch []contact.Contact, findNode FindNodeRPC) [][]contact.Contact {
	newBatch := [][]contact.Contact{batch}
	for i := 0; i < len(batch); i += environment.Alpha {
		var wg sync.WaitGroup
		for j := i; j < min((i+environment.Alpha), len(batch)); j++ {
			wg.Add(1)
			n := j
			go func() {
				defer wg.Done()
				kN, err := findNode(*contact.GetInstance(), batch[n], targetID)
				if err == nil {
					AddContact(batch[n])
					newBatch = append(newBatch, kN)
				}
			}()
		}
		wg.Wait()
	}
	return newBatch
}

// Merge a 2D slice to a 1D slice.
func mergeBatch(batch [][]contact.Contact) []contact.Contact {
	var mergedBatch []contact.Contact
	for i := 0; i < len(batch); i++ {
		mergedBatch = append(mergedBatch, batch[i]...)
	}
	return mergedBatch
}

// Remove duplicate Contacts, contacts with the same id, from batch.
func removeDuplicates(batch []contact.Contact) []contact.Contact {
	var cleanedBatch []contact.Contact
	for i := 0; i < len(batch); i++ {
		dupe := false
		currentID := batch[i].ID
		for j := 0; j < len(cleanedBatch); j++ {
			if currentID.Equals(cleanedBatch[j].ID) {
				dupe = true
			}
		}
		if !dupe {
			cleanedBatch = append(cleanedBatch, batch[i])
		}
	}
	return cleanedBatch
}

// Removes dead contacts by pinging and verifying if they are alive.
func removeDeadNodes(batch []contact.Contact, ping PingRpc) []contact.Contact {
	var wg sync.WaitGroup
	var deadNodes []int
	for i := 0; i < len(batch); i++ {
		wg.Add(1)
		n := i
		go func() {
			defer wg.Done()
			alive := ping(*contact.GetInstance(), batch[n])
			if alive {
				AddContact(batch[n])
			} else {
				deadNodes = append(deadNodes, n)
			}
		}()
	}

	wg.Wait()

	sort.Ints(deadNodes)

	for i := 0; i < len(deadNodes); i++ {
		idx := deadNodes[i] - i
		if idx != len(batch)-1 {
			batch = append(batch[:idx], batch[idx+1:]...)
		} else {
			batch = batch[:idx]
		}
	}

	return batch
}

// Resize batch to at most BucketSize in length.
func resize(batch []contact.Contact) []contact.Contact {
	if len(batch) > bucket.BucketSize {
		batch = batch[bucket.BucketSize:]
	}
	return batch
}

// Check if two slices of contacts are the same.
func isSame(batch []contact.Contact, newBatch []contact.Contact) bool {
	if len(batch) != len(newBatch) {
		return false
	}
	for i, c := range batch {
		if !c.ID.Equals(newBatch[i].ID) {
			return false
		}
	}
	return true
}
