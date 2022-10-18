package handler

import (
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/log"
	"D7024E/network/requestHandler"
	"D7024E/network/sender"
	"D7024E/node/kademlia/algorithms"
	"time"
)

// Depending on the RPC command initiate go routine.
func HandleCMD(msg []byte) {
	var rpcMessage rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(msg, &rpcMessage)

	go algorithms.AddContact(rpcMessage.Contact, rpc.Ping)

	log.INFO("OPERATION - [%s] SENDER - [%s]", rpcMessage.Cmd, rpcMessage.Contact.ID.String())
	startTime := time.Now()

	switch rpcMessage.Cmd {
	case "RESP":
		requestHandler.GetInstance().WriteRespone(rpcMessage.ReqID, msg)
	case "PING":
		rpc.Pong(rpcMessage.Contact, rpcMessage.ReqID, sender.UDPSender)
	case "RESH":
		rpc.RefreshResponse(rpcMessage.ID, rpcMessage.Contact, rpcMessage.ReqID, sender.UDPSender)
	case "STRE":
		rpc.StoreResponse(rpcMessage.Contact, rpcMessage.ReqID, rpcMessage.Content, sender.UDPSender)
	case "FINO":
		rpc.FindNodeResponse(rpcMessage.ReqID, rpcMessage.ID, rpcMessage.Contact, sender.UDPSender)
	case "FIVA":
		rpc.FindValueResponse(rpcMessage.Contact, rpcMessage.ReqID, rpcMessage.ID, sender.UDPSender)
	default:
		log.ERROR("UNKNOWN CMD")
	}
	log.INFO(
		"OPERATION - [%s] SENDER - [%s] DURATION - [%s]",
		rpcMessage.Cmd,
		rpcMessage.Contact.ID.String(),
		time.Since(startTime).String())
}
