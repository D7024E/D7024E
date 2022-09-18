package rpc

import (
	"net"
	"sync"

	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/log"
	"D7024E/network/requestHandler"
	"D7024E/network/sender"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/kademlia"
)

func FindNode(destNode id.KademliaID) {
	node := kademlia.GetInstance()
	rt := bucket.GetInstance()
	reqHandler := requestHandler.GetInstance()

	// Retrieves the alpha closest known contacts to the dest node and sends a findNode to each of them.
	// Each findNode is sent as a seperate goroutine, after a set amount of time any contact that has not responded
	// is assumed to be dead. A remove contact is sent for those contacts.
	// The bucket from the closest node that responded is then used to call a new findNode.
	alphaClosest := rt.FindClosestContacts(&destNode, node.Alpha)
	var alphaRes []contact.Contact

	var wg sync.WaitGroup
	for i := 0; i < node.Alpha; i++ {
		var reqID string = id.NewRandomKademliaID().String()
		var fino *[]byte
		rpc := rpcmarshal.RPC{
			Cmd:     "FINO",
			Sender:  node.Me,
			ReqID:   reqID,
			Content: destNode,
		}

		err := reqHandler.NewRequest(reqID)
		// If there is a request id collision, generate a new request id and try again.
		for err != nil {
			log.ERROR("%v", err)
			var reqID string = id.NewRandomKademliaID().String()
			rpc.ReqID = reqID
			err = reqHandler.NewRequest(reqID)
		}

		rpcmarshal.RpcMarshal(rpc, fino)
		// TODO: Retrieve port from environment.
		sender.UDPSender(net.IP(alphaClosest[0].Address), 4001, *fino)

		wg.Add(1)
		go reqHandler.ReadResponse(reqID)
	}
}
