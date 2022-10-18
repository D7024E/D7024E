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

// Initiate value refresh.
func NodeRefresh(value stored.Value) {
	stored.GetInstance().AddRefresh(value.ID)
	go NodeRefreshRec(value, rpc.RefreshRequest)
}

// Refresh a stored value in alpha closest nodes.
func NodeRefreshRec(value stored.Value, refresh RefreshRPC) bool {
	time.Sleep(time.Now().Add(500 * time.Millisecond).Sub(value.DeadAt))
	if !stored.GetInstance().IsRefreshed(value.ID) {
		return false
	} else {
		go NodeRefreshRec(value, refresh)
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
		NodeStore(value)
	}
	return true
}
