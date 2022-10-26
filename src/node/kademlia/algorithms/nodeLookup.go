package algorithms

import (
	"D7024E/environment"
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/network/sender"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"

	"sort"
	"sync"
)

type pingRPC func(contact.Contact, rpc.UDPSender) bool
type findNodeRPC func(contact.Contact, id.KademliaID, rpc.UDPSender) ([]contact.Contact, error)

// Node lookup initiator.
func NodeLookup(targetID id.KademliaID) []contact.Contact {
	batch := bucket.GetInstance().FindClosestContacts(&targetID, environment.Alpha)
	var nextBatch []contact.Contact
	for {
		nextBatch = NodeLookupRec(targetID, batch, rpc.FindNodeRequest, rpc.Ping)
		if isSame(batch, nextBatch) && len(nextBatch) >= 1 {
			return nextBatch
		} else {
			batch = nextBatch
		}
	}
}

// Algorithm for Node lookup.
func NodeLookupRec(targetID id.KademliaID, batch []contact.Contact, findNode findNodeRPC, ping pingRPC) []contact.Contact {
	if len(batch) == 0 {
		batch = append(batch, contact.Contact{ID: id.NewKademliaID("172.21.0.2"), Address: "172.21.0.2"})
	}
	rt := bucket.NewRoutingTable()
	var wg sync.WaitGroup
	for i := 0; i < len(batch); i += environment.Alpha {
		for j := i; j < min((i+environment.Alpha), len(batch)); j++ {
			wg.Add(1)
			n := j
			go func() {
				defer wg.Done()
				kN, err := findNode(batch[n], targetID, sender.UDPSender)
				if err == nil {
					for _, nodeContact := range kN {
						rt.AddContact(nodeContact)
					}
				}
			}()
		}
		wg.Wait()
	}
	return removeDeadNodes(rt.FindClosestContacts(&targetID, bucket.BucketSize), ping)
}

// min value of a and b.
func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

// Removes dead contacts by pinging and verifying if they are alive.
func removeDeadNodes(batch []contact.Contact, ping pingRPC) []contact.Contact {
	lock := sync.Mutex{}
	var deadNodes []int
	var wg sync.WaitGroup
	for i := 0; i < len(batch); i += environment.Alpha {
		for j := i; j < min((i+environment.Alpha), len(batch)); j++ {
			wg.Add(1)
			n := j
			go func() {
				defer wg.Done()
				alive := ping(batch[n], sender.UDPSender)
				if !alive {
					lock.Lock()
					defer lock.Unlock()
					deadNodes = append(deadNodes, n)
				}
			}()
		}
		wg.Wait()
	}

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
