package handler

import (
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/log"
	"D7024E/network/requestHandler"
	"D7024E/node/kademlia"
)

// Depending on the RPC command initiate go routine.
func HandleCMD(msg []byte) {
	var rpcMessage rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(msg, &rpcMessage)
	go kademlia.GetInstance().RoutingTable.AddContact(rpcMessage.Contact)
	switch rpcMessage.Cmd {
	case "RESP":
		// log.INFO("GOT RESPONSE")
		requestHandler.GetInstance().WriteRespone(rpcMessage.ReqID, msg)
	case "PING":
		rpc.Pong(*kademlia.GetInstance().Me, rpcMessage.Contact, rpcMessage.ReqID)
		// log.INFO("PONG DONE")
	case "STRE":
		rpc.StoreRequest(*kademlia.GetInstance().Me, rpcMessage.Contact, rpcMessage.Content)
		// log.INFO("STRE DONE")
	case "FINO":
		rpc.FindNodeRequest(*kademlia.GetInstance().Me, rpcMessage.Contact, rpcMessage.ID)
		// log.INFO("FINO DONE")
	case "FIVA":
		rpc.FindValueResponse(*kademlia.GetInstance().Me, rpcMessage.Contact, rpcMessage.ReqID, rpcMessage.ID)
		// log.INFO("FIVA DONE")
	default:
		log.ERROR("UNKNOWN CMD")
	}
}
