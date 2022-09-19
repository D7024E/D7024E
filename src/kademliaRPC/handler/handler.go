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
	case "PING":
		panic("help")
	case "PONG":
		panic("help")
	case "STRE":
<<<<<<< HEAD
		panic("help")
=======
		rpc.Store(rpcMessage.Content)
>>>>>>> 40246e51c7ac0c69dff5d951d075d2bfdf388c3f
	case "FINO":
		panic("help")
	case "FIVA":
		panic("help")
	default:
		log.ERROR("UNKNOWN CMD")
	}
}
