package algorithms

import (
	"D7024E/environment"
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/node/contact"
	"D7024E/node/kademlia"
	"D7024E/node/stored"
	"sync"
)

type storeRPC func(contact.Contact, contact.Contact, stored.Value) bool

// NodeStore value initiate.
func NodeStore(value stored.Value) bool {
	return KNodeStoreRec(value, rpc.StoreRequest)
}

// Store value in alpha nodes using fn.
func KNodeStoreRec(value stored.Value, fn storeRPC) bool {
	alphaClosest := NodeLookup(value.ID)
	if len(alphaClosest) > environment.Alpha {
		alphaClosest = alphaClosest[environment.Alpha:]
	}
	var wg sync.WaitGroup
	completed := true
	for _, c := range alphaClosest {
		wg.Add(1)
		target := c
		go func() {
			defer wg.Done()
			val := fn(*kademlia.GetInstance().Me, target, value)
			if !val {
				completed = val
			}
		}()
	}
	wg.Wait()
	if !completed {
		return KNodeStoreRec(value, fn)
	} else {
		return true
	}
}
