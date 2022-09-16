package node

import (
	err "D7024E/error"
)

type Kademlia struct {
	routingTable *RoutingTable
	values       Values
}

var KandemliaNode Kademlia

func init() {
	// Start routing table here
	// KandemliaNode.routingTable = NewRoutingTable(NewContact(NewKademliaID("STRING FOR ID"), "172.0.0.2"))
}

func (kademlia *Kademlia) LookupContact(target *Contact) {
	// TODO
}

func (kademlia *Kademlia) LookupValues(ID KademliaID) (Value, error) {
	for _, item := range kademlia.values {
		if item.ID == ID {
			return item, nil
		}
	}
	return Value{}, &err.ValueNotFound{}
}

func (kademlia *Kademlia) Store(value Value) {
	kademlia.values = append(kademlia.values, value)
}
