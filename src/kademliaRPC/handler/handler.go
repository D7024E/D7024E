package handler

import (
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/log"
	"D7024E/network/sender"
	"D7024E/node/kademlia/algorithms"
	"time"
)

// Depending on the RPC command initiate go routine.
func HandleCMD(msg []byte) []byte {
	var rpcMessage rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(msg, &rpcMessage)

	go algorithms.AddContact(rpcMessage.Contact, rpc.Ping)

	log.INFO("OPERATION - [%s] SENDER - [%s]", rpcMessage.Cmd, rpcMessage.Contact.ID.String())
	startTime := time.Now()

	var res []byte

	switch rpcMessage.Cmd {
	case "PING":
		res = rpc.Pong(rpcMessage.Contact)
	case "RESH":
		res = rpc.RefreshResponse(rpcMessage.ID, rpcMessage.Contact)
	case "STRE":
		res = rpc.StoreResponse(rpcMessage.Contact, rpcMessage.Content, sender.UDPSender)
	case "FINO":
		res = rpc.FindNodeResponse(rpcMessage.ID, rpcMessage.Contact)
	case "FIVA":
		res = rpc.FindValueResponse(rpcMessage.Contact, rpcMessage.ID, sender.UDPSender)
	default:
		log.ERROR("UNKNOWN CMD")
	}
	log.INFO(
		"OPERATION - [%s] SENDER - [%s] DURATION - [%s]",
		rpcMessage.Cmd,
		rpcMessage.Contact.ID.String(),
		time.Since(startTime).String())
	return res
}
