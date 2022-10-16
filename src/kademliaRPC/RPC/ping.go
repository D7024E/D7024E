package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/network/requestHandler"
	"D7024E/node/contact"
	"net"
)

// PING rpc.
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
	sender(ip, 4001, msg)
	err := requestHandler.GetInstance().ReadResponse(reqID, &msg)
	return !isError(err)
}

// PONG rpc.
// Ping target node, if there is a response return true, otherwise false.
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
