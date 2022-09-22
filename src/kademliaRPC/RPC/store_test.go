package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"testing"
	"time"
)

// Test if storeRequestMessage outputs correct message.
func TestStoreRequestMessageSuccess(t *testing.T) {
	rpc1 := rpcmarshal.RPC{
		Cmd: "STRE",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "THIS IS ADDRESS"},
		ReqID: newValidRequestID(),
		Content: stored.Value{
			Data:   "THIS IS DATA",
			ID:     *id.NewRandomKademliaID(),
			Ttl:    time.Second,
			DeadAt: time.Now().Add(time.Second),
		},
	}
	message := storeRequestMessage(rpc1.Contact, rpc1.ReqID, rpc1.Content)
	var rpc2 rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &rpc2)
	if !rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}

// Test if storeRequestMessage outputs correct message.
func TestStoreRequestMessageFail(t *testing.T) {
	rpc1 := rpcmarshal.RPC{
		Cmd: "STRE",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "THIS IS ADDRESS"},
		Content: stored.Value{
			Data:   "THIS IS DATA",
			ID:     *id.NewRandomKademliaID(),
			Ttl:    time.Second,
			DeadAt: time.Now().Add(time.Second),
		},
	}
	message := storeRequestMessage(rpc1.Contact, newValidRequestID(), rpc1.Content)
	var rpc2 rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &rpc2)
	if rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}

// Test if storeRespondMessage outputs correct message.
func TestStoreRespondMessageSuccess(t *testing.T) {
	rpc1 := rpcmarshal.RPC{
		Cmd: "RESP",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "THIS IS ADDRESS"},
		ReqID: newValidRequestID(),
	}
	message := storeRespondMessage(rpc1.Contact, rpc1.ReqID)
	var rpc2 rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &rpc2)
	if !rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}

// Test if storeRespondMessage outputs correct message.
func TestStoreRespondMessageFail(t *testing.T) {
	rpc1 := rpcmarshal.RPC{
		Cmd: "RESP",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "THIS IS ADDRESS"},
	}
	message := storeRespondMessage(rpc1.Contact, newValidRequestID())
	var rpc2 rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &rpc2)
	if rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}
