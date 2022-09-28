package algorithms

import (
	"D7024E/config"
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/log"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/kademlia"
	"D7024E/node/kademlia/kademliaSort"

	"sort"
	"sync"
)

func NodeLookup(destNode id.KademliaID) []contact.Contact {
	batch := kademlia.GetInstance().RoutingTable.FindClosestContacts(&destNode, config.Alpha)
	// batch = append(batch, kademlia.GetInstance().Me)
	return NodeLookupRec(destNode, batch)
}

func NodeLookupRec(destNode id.KademliaID, batch []contact.Contact) []contact.Contact {
	rt := bucket.GetInstance()
	me := rt.Me
	alpha := config.Alpha

	if len(batch) == 0 {
		batch = rt.FindClosestContacts(&destNode, alpha)
	}

	var newBatch [][]contact.Contact
	for i := 0; i < len(batch); i++ {
		var kN []contact.Contact
		kN, err := rpc.FindNodeRequest(me, batch[i], destNode)
		if err != nil {
			log.ERROR("%v", err)
		} else {
			newBatch = append(newBatch, kN)
		}
	}

	updatedBatch := mergeBatch(newBatch)
	updatedBatch = removeDuplicates(updatedBatch)
	updatedBatch = removeDeadNodes(updatedBatch, rpc.Ping)
	updatedBatch = getAllDistances(updatedBatch)
	updatedBatch = kademliaSort.SortContacts(updatedBatch)

	if len(updatedBatch) >= alpha {
		updatedBatch = updatedBatch[:alpha]
	}

	var sameBatch bool = true
	for i := 0; i < min(len(batch), len(updatedBatch)); i++ {
		if !batch[i].GetDistance().Equals(updatedBatch[i].GetDistance()) {
			sameBatch = false
		}
	}
	if sameBatch {
		return updatedBatch
	} else {
		return NodeLookupRec(destNode, updatedBatch)
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

// Note that this merge do not take duplicates into account.
func mergeBatch(batch [][]contact.Contact) []contact.Contact {
	var mergedBatch []contact.Contact
	for i := 0; i < len(batch); i++ {
		mergedBatch = append(mergedBatch, batch[i]...)
	}
	return mergedBatch
}

// Updates the distances in "batch" to be the distances to the current node then returns the new batch.
func getAllDistances(batch []contact.Contact) []contact.Contact {
	for i := 0; i < len(batch); i++ {
		relativeDistance := *batch[i].ID.CalcDistance(contact.GetInstance().ID)
		batch[i].SetDistance(&relativeDistance)
	}
	return batch
}

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

// Takes a list of contacts, and a function as arguments. The function should be Ping() or a test function.
func removeDeadNodes(batch []contact.Contact, fn func(contact.Contact, contact.Contact) bool) []contact.Contact {
	rt := bucket.GetInstance()
	me := rt.Me

	var deadNodes []int
	var wg sync.WaitGroup

	for i := 0; i < len(batch); i++ {
		wg.Add(1)
		n := i
		go func() {
			alive := fn(me, batch[n])
			if !alive {
				deadNodes = append(deadNodes, n)
			} else {
				AddContact(batch[n])
			}
			wg.Done()
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
