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

// Test Ping request to node that responds.
func TestPingSuccess(t *testing.T) {
	target := contact.Contact{
		Address: "0.0.0.0"}
	res := Ping(target, senderPingMockSuccess)
	if !res {
		t.FailNow()
	}
}

// Test Ping request to node that does not responds.
func TestPingFail(t *testing.T) {
	target := contact.Contact{
		Address: "0.0.0.0"}
	res := Ping(target, senderPingMockFail)
	if res {
		t.FailNow()
	}
}

// Test if pong does respond and that it returns correct response.
func TestPong(t *testing.T) {
	target := contact.Contact{
		Address: "0.0.0.0"}
	reqID := newValidRequestID()
	Pong(target, reqID, senderMockSuccess)

	var response []byte
	err := requestHandler.GetInstance().ReadResponse(reqID, &response)
	if err != nil {
		t.FailNow()
	}

	var rpcResponse rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(response, &rpcResponse)
	if !rpcResponse.Equals(&rpcmarshal.RPC{
		Cmd:     "RESP",
		Contact: *contact.GetInstance(),
		ReqID:   reqID}) {
		t.Errorf("wrong rpc response")
	}
}
