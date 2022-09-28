package kademlia

import (
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/stored"
	"net"
	"strings"
	"sync"
)

type KademliaNode struct {
	Alpha        int
	Me           contact.Contact
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
			instance.Me = *contact.GetInstance()
			instance.RoutingTable = bucket.GetInstance()
			instance.Values = stored.GetInstance()
		}
	}
	return instance
}

func (node *KademliaNode) LookupContact(target contact.Contact) {
	// closestContacts := node.RoutingTable.FindClosestContacts(target.ID, alpha)
	// TODO

}

func getAddress() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
