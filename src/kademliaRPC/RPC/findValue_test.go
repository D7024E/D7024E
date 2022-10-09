package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"errors"
	"testing"
	"time"
)

// Test FindValueRequestMessage output correct message.
func TestFindValueRequestMessageSuccess(t *testing.T) {
	rpc1 := rpcmarshal.RPC{
		Cmd: "FIVA",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "THIS IS ADDRESS"},
		ReqID: newValidRequestID(),
		ID:    *id.NewRandomKademliaID(),
	}
	message := findValueRequestMessage(rpc1.Contact, rpc1.ReqID, rpc1.ID)
	var rpc2 rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &rpc2)
	if !rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}

// Test FindValueRequestMessage output correct message.
func TestFindValueRequestMessageFail(t *testing.T) {
	rpc1 := rpcmarshal.RPC{
		Cmd: "FIVA",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "THIS IS ADDRESS"},
		ReqID: newValidRequestID(),
	}
	message := findValueRequestMessage(rpc1.Contact, rpc1.ReqID, *id.NewRandomKademliaID())
	var rpc2 rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &rpc2)
	if rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}

// Test successful return of request.
func TestFindValueRequestReturnSuccess(t *testing.T) {
	var message []byte
	rpcmarshal.RpcMarshal(rpcmarshal.RPC{Content: stored.Value{Data: "something"}}, &message)
	_, err := findValueRequestReturn(message, nil)
	if isError(err) {
		t.FailNow()
	}
}

// Test timeout error FindValueRequest.
func TestFindValueRequestReturnTimeout(t *testing.T) {
	_, err := findValueRequestReturn([]byte{}, errors.New("timeout"))
	if !isError(err) {
		t.FailNow()
	}
}

// Test not found error FindValueRequest.
func TestFindValueRequestReturnNotFound(t *testing.T) {
	var message []byte
	rpcmarshal.RpcMarshal(rpcmarshal.RPC{Content: stored.Value{}}, &message)
	_, err := findValueRequestReturn(message, nil)
	if !isError(err) {
		t.FailNow()
	}
}

// Test FindValueResponseMessage output in success case.
func TestFindValueResponseMessageSuccess(t *testing.T) {
	value := stored.Value{
		ID:  *id.NewRandomKademliaID(),
		Ttl: time.Hour,
	}
	err := stored.GetInstance().Store(value)
	if err != nil {
		t.FailNow()
	}
	rpc1 := rpcmarshal.RPC{
		Cmd: "RESP",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "THIS IS ADDRESS"},
		ReqID:   newValidRequestID(),
		Content: value,
	}
	message := findValueResponseMessage(rpc1.Contact, rpc1.ReqID, value.ID)
	var rpc2 rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &rpc2)
	if !rpc1.Content.Equals(&rpc2.Content) {
		t.FailNow()
	}
}

// Test FindValueResponseMessage output in not found case.
func TestFindValueResponseMessageNotFound(t *testing.T) {
	value := stored.Value{ID: *id.NewRandomKademliaID()}
	stored.GetInstance().Store(value)
	rpc1 := rpcmarshal.RPC{
		Cmd: "RESP",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "THIS IS ADDRESS"},
		ReqID:   newValidRequestID(),
		Content: value,
	}
	message := findValueResponseMessage(rpc1.Contact, rpc1.ReqID, *id.NewRandomKademliaID())
	var rpc2 rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &rpc2)
	if rpc1.Content.Equals(&rpc2.Content) {
		t.FailNow()
	}
}
