package algorithms

import (
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/kademlia"
	"D7024E/node/stored"
	"errors"
	"sync"
)

func NodeValueLookup(valueID id.KademliaID) (stored.Value, error) {
	closest := []contact.Contact{{Address: "172.21.0.2"}, {Address: "172.21.0.3"}, {Address: "172.21.0.4"}} // TODO NodeLookup(valueID)
	var wg sync.WaitGroup
	var result []stored.Value
	for _, c := range closest {
		wg.Add(1)
		target := c
		go func() {
			defer wg.Done()
			val, err := rpc.FindValueRequest(kademlia.GetInstance().Me, valueID, target)
			if err == nil {
				result = append(result, val)
			}
		}()
	}
	wg.Wait()
	for _, v := range result {
		if v.ID == valueID {
			return v, nil
		}
	}
	return stored.Value{}, errors.New("value not found")
}
