package rpc

import (
	"errors"
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

func FindNode(target contact.Contact, destNode id.KademliaID) (kNodes []contact.Contact, newErr error) {
	node := kademlia.GetInstance()
	reqHandler := requestHandler.GetInstance()

	var reqID string = id.NewRandomKademliaID().String()
	var fino []byte
	var response rpcmarshal.RPC

	rpc := rpcmarshal.RPC{
		Cmd:     "FINO",
		Contact: node.Me,
		ReqID:   reqID,
		ID:      destNode,
	}
	log.INFO("Creating find node RPC")

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
	log.INFO("Marshalling find node RPC")
	sender.UDPSender(net.IP(target.Address), config.Port, fino)
	log.INFO("Sending find node RPC")

	// Await and return the response.
	log.INFO("Waiting for find node response")
	err = reqHandler.ReadResponse(reqID, &fino)
	if err != nil {
		log.ERROR("%v", err)
		newErr := errors.New("No response")
		return nil, newErr
	}
	log.INFO("Received find node response")
	rpcmarshal.RpcUnmarshal(fino, &response)
	return response.KNodes, nil
}

// Creates a response RPC struct and populates it with the K, (K = 20), closest nodes to the destination node.
// Which is then sent back to the sender.
func RespondFindNode(rpc rpcmarshal.RPC) {
	node := kademlia.GetInstance()
	response := rpcmarshal.RPC{
		Cmd:     "RESP",
		Contact: node.Me,
		ReqID:   rpc.ReqID,
	}
	log.INFO("Creating new RPC-struct for find node")
	var marshaledResponse []byte
	target := &rpc.ID
	response.KNodes = bucket.GetInstance().FindClosestContacts(target, 20)
	log.INFO("Finding closest nodes to the find node target")
	rpcmarshal.RpcMarshal(response, &marshaledResponse)
	log.INFO("Marshalling find node RPC-response")
	sender.UDPSender(net.IP(rpc.Contact.Address), config.Port, marshaledResponse)
	log.INFO("Sending find node response")
}
