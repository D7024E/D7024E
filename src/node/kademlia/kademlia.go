package kademlia

import (
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/stored"
	"sync"
)

type KademliaNode struct {
	Alpha        int
	Me           *contact.Contact
	RoutingTable *bucket.RoutingTable
	Values       *stored.Stored
}

const alpha = 3

var lock = &sync.Mutex{}
var instance *KademliaNode

func GetInstance() *KademliaNode {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &KademliaNode{}
			instance.Me = contact.GetInstance()
			instance.RoutingTable = bucket.GetInstance()
			instance.Values = stored.GetInstance()
		}
	}
	return instance
}
