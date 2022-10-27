package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/node/contact"
	"D7024E/node/id"
	"net"
	"testing"
)

// UDPSender mockup that simulates a successful response.
func senderPingMockSuccess(_ net.IP, _ int, message []byte) ([]byte, error) {
	var request rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &request)

	var response []byte
	rpcmarshal.RpcMarshal(
		rpcmarshal.RPC{
			Cmd:     "RESP",
			Contact: *contact.GetInstance(),
		},
		&response)
	return response, nil
}

// UDPSender mockup that simulates no response.
func senderPingMockFail(_ net.IP, _ int, _ []byte) ([]byte, error) {
	return nil, nil
}

// Test Ping request to node that responds.
func TestPingSuccess(t *testing.T) {
	target := contact.Contact{
		ID:      id.NewRandomKademliaID(),
		Address: "0.0.0.0"}
	res := Ping(target, senderPingMockSuccess)
	if !res {
		t.Fatalf("retrieved no response, when given a response")
	}
}

// Test Ping request to node that does not responds.
func TestPingFail(t *testing.T) {
	target := contact.Contact{
		ID:      id.NewRandomKademliaID(),
		Address: "0.0.0.0"}
	res := Ping(target, senderPingMockFail)
	if res {
		t.Fatalf("retrieved response, when no response given")
	}
}

// Test if pong does respond and that it returns correct response.
func TestPong(t *testing.T) {
	target := contact.Contact{Address: "0.0.0.0"}
	response := Pong(target)
	var rpcResponse rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(response, &rpcResponse)
	if !rpcResponse.Equals(&rpcmarshal.RPC{
		Cmd:     "RESP",
		Contact: *contact.GetInstance(),
	}) {
		t.Fatalf("wrong rpc response")
	}
}
