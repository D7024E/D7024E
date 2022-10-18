package algorithms

import (
	"D7024E/environment"
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/network/sender"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"fmt"
	"sync"
	"time"
)

type RefreshRPC func(id.KademliaID, contact.Contact, rpc.UDPSender) bool
type NodeStoreAlgorithm func(value stored.Value) bool

// Initiate value refresh.
func NodeRefresh(value stored.Value) {
	stored.GetInstance().AddRefresh(value.ID)
	go NodeRefreshRec(value, rpc.RefreshRequest, NodeStore)
}

// Refresh a stored value in alpha closest nodes.
func NodeRefreshRec(value stored.Value, refresh RefreshRPC, store NodeStoreAlgorithm) bool {
	if !stored.GetInstance().IsRefreshed(value.ID) {
		fmt.Println("No longer refreshing")
		return false
	} else {
		go func() {
			time.Sleep(value.Ttl)
			NodeRefreshRec(value, refresh, store)
		}()
	}
	alphaClosest := NodeLookup(value.ID)
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
			val := refresh(value.ID, target, sender.UDPSender)
			if !val {
				completed = val
			}
		}()
	}
	wg.Wait()
	if !completed {
		fmt.Println("Storing for refresh")
		store(value)
	}
	return true
}
