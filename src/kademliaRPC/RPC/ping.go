package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/node/contact"
)

// Ping target node, if there is a response return true, otherwise false.
func Ping(target contact.Contact, sender UDPSender) bool {
	var msg []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "PING",
			Contact: *contact.GetInstance(),
		},
		&msg)

	resMessage, err := sender(parseIP(target.Address), 4001, msg)
	if isError(err) || resMessage == nil {
		return false
	}

	go AddContact(target, Ping)

	var rpcMessage rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(resMessage, &rpcMessage)
	return true
}

// Respond to ping.
func Pong(target contact.Contact) []byte {
	var msg []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
		},
		&msg)

	return msg
}
