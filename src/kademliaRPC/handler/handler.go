package handler

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/log"
)

// Depending on the RPC command initiate go routine.
func HandleCMD(msg []byte) {
	var rpcMessage rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(msg, &rpcMessage)
	switch rpcMessage.Cmd {
	case "RESP":
		panic("help")
	case "PING":
		panic("help")
	case "STRE":
		panic("help")
	case "FINO":
		panic("help")
	case "FIVA":
		panic("help")
	default:
		log.ERROR("UNKNOWN CMD")
	}
}
