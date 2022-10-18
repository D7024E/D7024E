package algorithms

import (
	"D7024E/environment"
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/network/sender"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"sync"
)

type lookupAlgorithm func(id.KademliaID) []contact.Contact
type storeRPC func(contact.Contact, stored.Value, rpc.UDPSender) bool

// NodeStore value initiate.
func NodeStore(value stored.Value) bool {
	return AlphaNodeStoreRec(value, rpc.StoreRequest, NodeLookup)
}

// Store value in alpha nodes using fn.
func AlphaNodeStoreRec(value stored.Value, store storeRPC, lookup lookupAlgorithm) bool {
	alphaClosest := lookup(value.ID)
	if len(alphaClosest) > environment.Alpha {
		alphaClosest = alphaClosest[:environment.Alpha]
	}
	var wg sync.WaitGroup
	completed := true
	for _, c := range alphaClosest {
		wg.Add(1)
		target := c
		go func() {
			defer wg.Done()
			val := store(target, value, sender.UDPSender)
			if !val {
				completed = val
			}
		}()
	}
	wg.Wait()
	if !completed {
		return AlphaNodeStoreRec(value, store, lookup)
	} else {
		return true
	}
}
