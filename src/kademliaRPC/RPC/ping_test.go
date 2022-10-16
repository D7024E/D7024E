package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/network/requestHandler"
	"D7024E/node/contact"
	"net"
	"testing"
)

// UDPSender mockup that simulates a successful response.
func senderPingMockSuccess(_ net.IP, _ int, message []byte) {
	var request rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &request)

	var response []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
			ReqID:   request.ReqID,
		},
		&response)

	requestHandler.GetInstance().WriteRespone(
		request.ReqID,
		response)
}

// UDPSender mockup that simulates no response.
func senderPingMockFail(_ net.IP, _ int, _ []byte) {}

func TestPingSuccess(t *testing.T) {
	target := contact.Contact{
		Address: "0.0.0.0"}
	res := Ping(target, senderPingMockSuccess)
	if !res {
		t.FailNow()
	}
}

func TestPingFail(t *testing.T) {
	target := contact.Contact{
		Address: "0.0.0.0"}
	res := Ping(target, senderPingMockFail)
	if res {
		t.FailNow()
	}
}
