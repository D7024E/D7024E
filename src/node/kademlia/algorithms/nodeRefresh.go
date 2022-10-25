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
	go func() {
		deadline := time.Now().Add(value.Ttl - 5*time.Second)
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
		fmt.Println("[NODE REFRESH] - stopped refresh for: ", value.ID.String())
		return false
	} else {
		go func() {
			deadline := time.Now().Add(value.Ttl - 5*time.Second)
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
	var failed []string
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
				failed = append(failed, target.Address)
			}
		}()
	}
	wg.Wait()
	fmt.Println(failed)
	fmt.Println("[NODE REFRESH] - sent refresh for value: ", value.ID.String())
	if restore {
		fmt.Println("[NODE REFRESH] - store value again, for value: ", value.ID.String())
		go NodeStore(value)
	}
	return true
}
