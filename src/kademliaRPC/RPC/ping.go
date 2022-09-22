package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/network/requestHandler"
	"D7024E/network/sender"
	"D7024E/node/contact"
	"net"
)

// PING rpc.
// Ping target node, if there is a response return true, otherwise false.
func Ping(me contact.Contact, target contact.Contact) bool {
	reqID := newValidRequestID()
	msg := PingMessage(me, reqID)
	ip := net.ParseIP(target.Address)
	sender.UDPSender(ip, 4001, msg)
	err := requestHandler.GetInstance().ReadResponse(reqID, &msg)
	return !isError(err)
}

func PingMessage(me contact.Contact, reqID string) []byte {

	var msg []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "PING",
			Contact: me,
			ReqID:   reqID,
		}, &msg)
	return msg
}

// PONG rpc.
// Ping target node, if there is a response return true, otherwise false.
func Pong(me contact.Contact, target contact.Contact, reqID string) {

	msg := PongMessage(me, reqID)
	sender.UDPSender(net.ParseIP(target.Address), 4001, msg)
}

func PongMessage(me contact.Contact, reqID string) []byte {

	var msg []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: me,
			ReqID:   reqID,
		}, &msg)
	return msg
}
