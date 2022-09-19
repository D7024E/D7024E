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

func ping(target contact.Contact) {

	node := kademlia.GetInstance()
	var reqID string = id.NewRandomKademliaID().String()
	reqTable := requestHandler.GetInstance()
	err := reqTable.NewRequest(reqID)

	rpc := rpcmarshal.RPC{
		Cmd:     "PING",
		Contact: node.Me,
		ReqID:   reqID,
	}

	if err != nil {
		log.ERROR("%v", err)
		var reqID string = id.NewRandomKademliaID().String()
		rpc.ReqID = reqID
		err = reqTable.NewRequest(reqID)
	}

	var msg *[]byte
	rpcmarshal.RpcMarshal(rpc, msg)

	str := target.Address

	ip := net.ParseIP(str)

	sender.UDPSender(ip, 4001, "message")

}
