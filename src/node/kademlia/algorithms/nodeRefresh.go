package algorithms

import (
	"D7024E/environment"
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/network/sender"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"sync"
	"time"
)

type RefreshRPC func(id.KademliaID, contact.Contact, rpc.UDPSender) bool
type NodeStoreAlgorithm func(value stored.Value) bool

// Initiate value refresh.
func NodeRefresh(value stored.Value) {
	go func() {
		alphaClosest := NodeLookup(value.ID)
		if len(alphaClosest) > environment.Alpha {
			alphaClosest = alphaClosest[:environment.Alpha]
		}
		NodeRefreshRec(value, alphaClosest, rpc.RefreshRequest)
	}()
	stored.GetInstance().AddRefresh(value.ID)
}

// Refresh a stored value in alpha closest nodes.
func NodeRefreshRec(value stored.Value, alphaClosest []contact.Contact, refresh RefreshRPC) bool {
	if !stored.GetInstance().IsRefreshed(value.ID) {
		return false
	} else {
		go func() {
			time.Sleep(value.Ttl / 2)
			NodeRefreshRec(value, alphaClosest, refresh)
		}()
	}
	var wg sync.WaitGroup
	for _, c := range alphaClosest {
		wg.Add(1)
		target := c
		go func() {
			defer wg.Done()
			refresh(value.ID, target, sender.UDPSender)
		}()
	}
	wg.Wait()
	return true
}
