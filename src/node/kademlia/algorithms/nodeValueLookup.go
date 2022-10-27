package algorithms

import (
	"D7024E/environment"
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/log"
	"D7024E/network/sender"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"errors"
	"sync"
)

type findValueRPC func(id.KademliaID, contact.Contact, rpc.UDPSender) (stored.Value, error)

// Initiates NodeValueLookup with alpha nodes.
func NodeValueLookup(valueID id.KademliaID) (stored.Value, error) {
	alphaClosest := NodeLookup(valueID)
	kAddress := []string{}
	for _, c := range alphaClosest {
		kAddress = append(kAddress, c.Address)
	}
	if len(alphaClosest) > environment.Alpha {
		alphaClosest = alphaClosest[:environment.Alpha]
	}
	return alphaNodeValueLookup(valueID, alphaClosest, rpc.FindValueRequest)
}

// Lookup value with valueID in alpha closest nodes using fn.
func alphaNodeValueLookup(valueID id.KademliaID, alphaClosest []contact.Contact, fn findValueRPC) (stored.Value, error) {
	log.INFO("[NODE VALUE LOOKUP] - lookup value with id: ", valueID.String(), " in \n", alphaClosest)
	var wg sync.WaitGroup
	var result []stored.Value
	lock := sync.Mutex{}
	for _, c := range alphaClosest {
		wg.Add(1)
		target := c
		go func() {
			defer wg.Done()
			val, err := fn(valueID, target, sender.UDPSender)
			if err == nil {
				lock.Lock()
				defer lock.Unlock()
				result = append(result, val)
			}
		}()
	}
	wg.Wait()
	for _, v := range result {
		if id.NewKademliaID(v.Data).Equals(&valueID) {
			return v, nil
		}
	}
	return stored.Value{}, errors.New("value not found")
}
