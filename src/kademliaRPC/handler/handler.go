package handler

import (
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/log"
)

// Depending on the RPC command initiate go routine.
func HandleCMD(msg []byte) {
	var rpcMessage rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(msg, &rpcMessage)
	switch rpcMessage.Cmd {
	case "PING":
		panic("help")
	case "PONG":
		panic("help")
	case "STRE":
		panic("help")
	case "FINO":
		go rpc.RespondFindNode(rpcMessage)
	case "FIVA":
		panic("help")
	default:
		log.ERROR("UNKNOWN CMD")
	}
}
