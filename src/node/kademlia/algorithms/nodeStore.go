package algorithms

import (
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/node/contact"
	"D7024E/node/kademlia"
	"D7024E/node/stored"
	"sync"
)

type storeRPC func(contact.Contact, contact.Contact, stored.Value) bool

func NodeStore(value stored.Value) bool {
	return AlphaNodeStoreRec(value, rpc.StoreRequest)
}

func AlphaNodeStoreRec(value stored.Value, fn storeRPC) bool {
	alphaClosest := []contact.Contact{{Address: "172.21.0.2"}, {Address: "172.21.0.3"}, {Address: "172.21.0.4"}} // TODO NodeLookup(valueID)
	var wg sync.WaitGroup
	completed := true
	for _, c := range alphaClosest {
		wg.Add(1)
		target := c
		go func() {
			defer wg.Done()
			val := fn(kademlia.GetInstance().Me, target, value)
			if !val {
				completed = val
			}
		}()
	}
	wg.Wait()
	if !completed {
		return AlphaNodeStoreRec(value, fn)
	} else {
		return true
	}
}
