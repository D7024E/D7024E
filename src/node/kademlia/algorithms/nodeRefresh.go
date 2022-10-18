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
	go func() {
		time.Sleep(value.Ttl / 2)
		NodeRefreshRec(value, rpc.RefreshRequest, NodeStore)
	}()
	stored.GetInstance().AddRefresh(value.ID)
}

// Refresh a stored value in alpha closest nodes.
func NodeRefreshRec(value stored.Value, refresh RefreshRPC, store NodeStoreAlgorithm) bool {
	if !stored.GetInstance().IsRefreshed(value.ID) {
		return false
	} else {
		go func() {
			time.Sleep(value.Ttl / 2)
			NodeRefreshRec(value, refresh, store)
		}()
	}
	fmt.Println("Refresh")
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
			res := refresh(value.ID, target, sender.UDPSender)
			if !res {
				fmt.Println(target.Address)
				completed = res
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
