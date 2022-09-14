package node

import err "D7024E/error"

type Kademlia struct {
	routingTable *RoutingTable
	objects      Objects
}

var KandemliaNode Kademlia

func init() {
	KandemliaNode.routingTable = NewRoutingTable(NewContact(NewKademliaID("STRING FOR ID"), "172.0.0.2"))
}

func (kademlia *Kademlia) LookupContact(target *Contact) {
	// TODO
}

func (kademlia *Kademlia) LookupObject(hash string) (Object, error) {
	for _, item := range kademlia.objects {
		if item.Hash == hash {
			return item, nil
		}
	}
	return Object{}, &err.ObjectNotFound{}
}

func (kademlia *Kademlia) Store(object Object) {
	// TODO
}
