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

type findValueRPC func(contact.Contact, id.KademliaID, contact.Contact) (stored.Value, error)

func NodeValueLookup(valueID id.KademliaID) (stored.Value, error) {
	alphaClosest := []contact.Contact{{Address: "172.21.0.2"}, {Address: "172.21.0.3"}, {Address: "172.21.0.4"}} // TODO NodeLookup(valueID)
	return AlphaNodeValueLookup(valueID, alphaClosest, rpc.FindValueRequest)
}

func AlphaNodeValueLookup(valueID id.KademliaID, alphaClosest []contact.Contact, fn findValueRPC) (stored.Value, error) {
	var wg sync.WaitGroup
	var result []stored.Value
	for _, c := range alphaClosest {
		wg.Add(1)
		target := c
		go func() {
			defer wg.Done()
			val, err := fn(kademlia.GetInstance().Me, valueID, target)
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
