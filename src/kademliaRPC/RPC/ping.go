package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/log"
	"D7024E/network/requestHandler"
	"D7024E/network/sender"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/kademlia"
	"net"
)

// ping node
func ping(contact contact.Contact) {

	node := kademlia.GetInstance()
	//request new request id
	var reqID string = id.NewRandomKademliaID().String()
	reqTable := requestHandler.GetInstance()
	err := reqTable.NewRequest(reqID)

	rpc := rpcmarshal.RPC{
		Cmd:     "PING",
		Contact: node.Me,
		ReqID:   reqID,
	}

	// if reqId already exists
	if err != nil {
		log.ERROR("%v", err)
		var reqID string = id.NewRandomKademliaID().String()
		rpc.ReqID = reqID
		err = reqTable.NewRequest(reqID)
	}

	var msg *[]byte
	rpcmarshal.RpcMarshal(rpc, msg)

	ip := net.ParseIP(contact.Address)

	sender.UDPSender(ip, 4001, *msg)

}

// pong response responds to the request id from ping
func pongResponse(contact contact.Contact, kademliaID id.KademliaID) {

	node := kademlia.GetInstance()
	var reqID string = kademliaID.String()

	rpc := rpcmarshal.RPC{
		Cmd:     "PONG",
		Contact: node.Me,
		ReqID:   reqID,
	}

	var msg *[]byte
	rpcmarshal.RpcMarshal(rpc, msg)

	ip := net.ParseIP(contact.Address)

	sender.UDPSender(ip, 4001, *msg)
}
