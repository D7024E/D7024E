package algorithms

import (
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/node/contact"
	"D7024E/node/kademlia"
	"D7024E/node/stored"
	"sync"
)

func NodeStore(value stored.Value) bool {
	closest := []contact.Contact{{Address: "172.21.0.2"}, {Address: "172.21.0.3"}, {Address: "172.21.0.4"}} // TODO NodeLookup(valueID)
	var wg sync.WaitGroup
	completed := true
	for _, c := range closest {
		wg.Add(1)
		target := c
		go func() {
			defer wg.Done()
			val := rpc.StoreRequest(*kademlia.GetInstance().Me, target, value)
			if !val {
				completed = val
			}
		}()
	}
	wg.Wait()
	if !completed {
		return NodeStore(value)
	} else {
		return true
	}
}
