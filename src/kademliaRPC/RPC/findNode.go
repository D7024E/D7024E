package rpc

import (
	"net"

	"D7024E/config"
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/log"
	"D7024E/network/requestHandler"
	"D7024E/network/sender"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/kademlia"
)

func FindNode(target contact.Contact, destNode id.KademliaID) (kNodes []contact.Contact) {
	node := kademlia.GetInstance()
	reqHandler := requestHandler.GetInstance()

	var reqID string = id.NewRandomKademliaID().String()
	var fino []byte
	var response rpcmarshal.RPC

	rpc := rpcmarshal.RPC{
		Cmd:         "FINO",
		Sender:      node.Me,
		ReqID:       reqID,
		Destination: destNode,
	}

	// Attempt to lock th request id in the reqTable.
	err := reqHandler.NewRequest(reqID)
	// If there is a request id collision, generate a new request id and try again.
	for err != nil {
		log.ERROR("%v", err)
		var reqID string = id.NewRandomKademliaID().String()
		rpc.ReqID = reqID
		err = reqHandler.NewRequest(reqID)
	}

	// Marshal the rpc struct and send it to the target.
	rpcmarshal.RpcMarshal(rpc, &fino)
	sender.UDPSender(net.IP(target.Address), config.Port, fino)

	// Await and return the response.
	reqHandler.ReadResponse(reqID, &fino)
	rpcmarshal.RpcUnmarshal(fino, &response)
	return response.KNodes
}

// Creates a response RPC struct and populates it with the K, (K = 20), closest nodes to the destination node.
// Which is then sent back to the sender.
func RespondFindNode(rpc rpcmarshal.RPC) {
	node := kademlia.GetInstance()
	response := rpcmarshal.RPC{
		Cmd:    "RESP",
		Sender: node.Me,
		ReqID:  rpc.ReqID,
	}
	var marshaledResponse []byte
	target := &rpc.Destination
	response.KNodes = bucket.GetInstance().FindClosestContacts(target, 20)
	rpcmarshal.RpcMarshal(response, &marshaledResponse)
	sender.UDPSender(net.IP(rpc.Sender.Address), config.Port, marshaledResponse)
}
