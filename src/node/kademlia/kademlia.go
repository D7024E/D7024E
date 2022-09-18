package kademlia

import (
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/stored"
	"sync"
)

type KademliaNode struct {
	Alpha        int
	Me           contact.Contact
	RoutingTable *bucket.RoutingTable
	Values       *stored.Stored
}

var lock = &sync.Mutex{}
var instance *KademliaNode

func GetInstance() *KademliaNode {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &KademliaNode{}
		}
	}
	return instance
}

// instance := kademlia.GetInstance()
// instance.Me to get Me Contact
// instance.RoutingTable to get routing table
// instance.Values to get the stored values
