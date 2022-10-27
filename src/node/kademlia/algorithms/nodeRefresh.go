package algorithms

import (
	"D7024E/environment"
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/log"
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
	stored.GetInstance().AddRefresh(value.ID)
	go func() {
		deadline := time.Now().Add(value.TTL - 5*time.Second)
		alphaClosest := NodeLookup(value.ID)
		if len(alphaClosest) > environment.Alpha {
			alphaClosest = alphaClosest[:environment.Alpha]
		}
		time.Sleep(time.Until(deadline))
		NodeRefreshRec(value, alphaClosest, rpc.RefreshRequest)
	}()
}

// Refresh a stored value in alpha closest nodes.
func NodeRefreshRec(value stored.Value, alphaClosest []contact.Contact, refresh RefreshRPC) bool {
	if !stored.GetInstance().IsRefreshed(value.ID) {
		log.INFO("[NODE REFRESH] - stopped refresh for: %v", value.ID.String())
		return false
	} else {
		go func() {
			deadline := time.Now().Add(value.TTL - 5*time.Second)
			alphaClosest := NodeLookup(value.ID)
			if len(alphaClosest) > environment.Alpha {
				alphaClosest = alphaClosest[:environment.Alpha]
			}
			time.Sleep(time.Until(deadline))
			NodeRefreshRec(value, alphaClosest, refresh)
		}()
	}

	lock := sync.Mutex{}
	restore := false
	var wg sync.WaitGroup
	for _, c := range alphaClosest {
		wg.Add(1)
		target := c
		go func() {
			defer wg.Done()
			refreshed := refresh(value.ID, target, sender.UDPSender)
			if !refreshed {
				lock.Lock()
				defer lock.Unlock()
				restore = true
			}
		}()
	}
	wg.Wait()
	log.INFO("[NODE REFRESH] - sent refresh for value: %v", value.ID.String())
	if restore {
		log.INFO("[NODE REFRESH] - store value again, for value: %v", value.ID.String())
		go NodeStore(value)
	}
	return true
}
