package rpc

import (
	"D7024E/node/bucket"
	"D7024E/node/id"
	"D7024E/node/kademlia"
)

func FindNode(destNode id.KademliaID) {
	instance := kademlia.GetInstance()
	me := instance.Me.ID
	rt := bucket.GetInstance()

	contact := rt.FindClosestContacts(&destNode, 1)

}
