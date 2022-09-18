package rpc

import (
	"net"

	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/network/sender"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/kademlia"
)

func FindNode(destNode id.KademliaID) {
	node := kademlia.GetInstance()
	rt := bucket.GetInstance()

	// Retrieves the alpha closest known contacts to the dest node and sends a findNode to each of them.
	// Each findNode is sent as a seperate goroutine, after a set amount of time any contact that has not responded
	// is assumed to be dead. A remove contact is sent for those contacts.
	// The bucket from the closest node that responded is then used to call a new findNode.
	alphaClosest := rt.FindClosestContacts(&destNode, node.Alpha)
	var alphaRes []contact.Contact

	for i := 0; i < node.Alpha; i++ {
		var fino *[]byte
		rpc := rpcmarshal.RPC{
			Cmd:     "FINO",
			Sender:  node.Me,
			Content: destNode,
		}

		rpcmarshal.RpcMarshal(rpc, fino)
		sender.UDPSender(net.IP(alphaClosest[0].Address), 4001, *fino)
	}
}

func alphaSender(node *kademlia.KademliaNode, destNode id.KademliaID, query contact.Contact, res *[]contact.Contact) {
	var fino *[]byte
	rpc := rpcmarshal.RPC{
		Cmd:     "FINO",
		Sender:  node.Me,
		Content: destNode,
	}
	rpcmarshal.RpcMarshal(rpc, fino)
	sender.UDPSender(net.IP(query.Address), 4001, *fino)
}
