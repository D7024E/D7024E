package algorithms

import (
	"D7024E/environment"
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"errors"
	"sync"
)

type findValueRPC func(contact.Contact, id.KademliaID, contact.Contact) (stored.Value, error)

// Initiates NodeValueLookup with alpha nodes.
func NodeValueLookup(valueID id.KademliaID) (stored.Value, error) {
	alphaClosest := NodeLookup(valueID)
	if len(alphaClosest) > environment.Alpha {
		alphaClosest = alphaClosest[:environment.Alpha]
	}
	return alphaNodeValueLookup(valueID, alphaClosest, rpc.FindValueRequest)
}

// Lookup value with valueID in alpha closest nodes using fn.
func alphaNodeValueLookup(valueID id.KademliaID, alphaClosest []contact.Contact, fn findValueRPC) (stored.Value, error) {
	var wg sync.WaitGroup
	var result []stored.Value
	for _, c := range alphaClosest {
		wg.Add(1)
		target := c
		go func() {
			defer wg.Done()
			val, err := fn(*contact.GetInstance(), valueID, target)
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
