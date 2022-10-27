package algorithms

import (
	"D7024E/environment"
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/network/sender"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"fmt"

	"sync"
)

type pingRPC func(contact.Contact, rpc.UDPSender) bool
type findNodeRPC func(contact.Contact, id.KademliaID, rpc.UDPSender) ([]contact.Contact, error)

// Node lookup initiator.
func NodeLookup(targetID id.KademliaID) []contact.Contact {
	rt := bucket.NewRoutingTable()
	rt.AddContact(contact.Contact{ID: id.NewKademliaID("172.21.0.2"), Address: "172.21.0.2"})
	batch := bucket.GetInstance().FindClosestContacts(&targetID, environment.Alpha)
	for {
		fmt.Println(batch)
		if isSame(rt.FindClosestContacts(&targetID, bucket.BucketSize), batch) && len(batch) >= 2 {
			return batch
		} else {
			batch = nodeLookup(targetID, rt, rpc.Ping, rpc.FindNodeRequest)
		}

	}
}

// Add multiple contacts to routing table if they respond to a ping.
func AddContacts(rt *bucket.RoutingTable, batch []contact.Contact, ping pingRPC) {
	var wg sync.WaitGroup
	for i := 0; i < len(batch); i += environment.Alpha {
		for j := 0; j < min(len(batch), environment.Alpha); j++ {
			wg.Add(1)
			go func(target contact.Contact) {
				defer wg.Done()
				if !(rt.FindClosestContacts(target.ID, 1)[0].ID).Equals(target.ID) {
					if rpc.Ping(target, sender.UDPSender) {
						rt.AddContact(target)
					}
				}
			}(batch[j])
		}
		wg.Wait()
	}
}

// Algorithm for Node lookup.
func nodeLookup(targetID id.KademliaID, rt *bucket.RoutingTable, ping pingRPC, findNode findNodeRPC) []contact.Contact {
	batch := rt.FindClosestContacts(&targetID, bucket.BucketSize)
	var wg sync.WaitGroup
	for i := 0; i < len(batch); i += environment.Alpha {
		for j := i; j < min((i+environment.Alpha), len(batch)); j++ {
			wg.Add(1)
			target := batch[j]
			go func() {
				defer wg.Done()
				kN, err := findNode(target, targetID, sender.UDPSender)
				if err == nil {
					rt.AddContact(target)
					AddContacts(rt, kN, ping)
				}
			}()
		}
		wg.Wait()
	}
	return rt.FindClosestContacts(&targetID, bucket.BucketSize)
}

// min value of a and b.
func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
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
