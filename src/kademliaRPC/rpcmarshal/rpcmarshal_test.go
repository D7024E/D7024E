package rpcmarshal

import (
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"testing"
	"time"
)

func TestRpcMarshal(t *testing.T) {
	rpc := RPC{
		Cmd: "THIS IS CMD",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "0.0.0.0",
		},
		ReqID: "THIS IS REQUEST ID",
		ID:    *id.NewRandomKademliaID(),
		Content: stored.Value{
			Data:   "THIS IS DATA",
			ID:     *id.NewRandomKademliaID(),
			Ttl:    time.Second,
			DeadAt: time.Now().Add(time.Second),
		},
	}

	var message []byte
	var rpc2 RPC
	err := RpcMarshal(rpc, &message)
	if err != nil {
		t.FailNow()
	}
	RpcUnmarshal(message, &rpc2)

	if !rpc.Equals(&rpc2) {
		t.FailNow()
	}
}
