package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/node/contact"
	"D7024E/node/id"
	"testing"
)

func TestPingRequestMessageSuccess(t *testing.T) {
	rpc1 := rpcmarshal.RPC{
		Cmd: "PING",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "THIS IS ADDRESS"},
		ReqID: newValidRequestID(),
	}
	msg := PingMessage(rpc1.Contact, rpc1.ReqID)
	var rpc2 rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(msg, &rpc2)
	if !rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}

func TestPingRequestMessageFail(t *testing.T) {
	rpc1 := rpcmarshal.RPC{
		Cmd: "PING",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "THIS IS ADDRESS"},
	}
	msg := PingMessage(rpc1.Contact, newValidRequestID())
	var rpc2 rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(msg, &rpc2)
	if rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}

func TestPongRequestMessageSuccess(t *testing.T) {
	rpc1 := rpcmarshal.RPC{
		Cmd: "RESP",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "THIS IS ADDRESS"},
		ReqID: newValidRequestID(),
	}
	msg := PongMessage(rpc1.Contact, rpc1.ReqID)
	var rpc2 rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(msg, &rpc2)
	if !rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}

func TestPongRequestMessageFail(t *testing.T) {
	rpc1 := rpcmarshal.RPC{
		Cmd: "RESP",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "THIS IS ADDRESS"},
	}
	msg := PongMessage(rpc1.Contact, newValidRequestID())
	var rpc2 rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(msg, &rpc2)
	if rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}
