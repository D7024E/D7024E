package kademlia_rpc

import "D7024E/log"

// Depending on the RPC command initiate go routine.
func InitiateCMD(msg []byte) {
	var rpc RPC
	RpcUnmarshal(msg, &rpc)
	switch rpc.Cmd {
	case "PING":
		panic("help")
	case "PONG":
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
