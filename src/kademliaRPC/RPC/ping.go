package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/log"
	"D7024E/node/contact"
	"net"
)

// Ping target node, if there is a response return true, otherwise false.
func Ping(target contact.Contact, sender UDPSender) bool {
	reqID := newValidRequestID()
	var msg []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "PING",
			Contact: *contact.GetInstance(),
			ReqID:   reqID,
		},
		&msg)
	ip := net.ParseIP(target.Address)
	resMessage, err := sender(ip, 4001, msg)
	if err != nil {
		log.ERROR("Error when sending rpc")
		return false
	}
	var rpcMessage rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(resMessage, &rpcMessage)
	return !isError(err)
}

// Respond to ping.
func Pong(target contact.Contact, reqID string, sender UDPSender) {
	var msg []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
			ReqID:   reqID,
		},
		&msg)
	sender(net.ParseIP(target.Address), 4001, msg)
}
