package rpc

import (
	"D7024E/kademliaRPC/rpcmarshal"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"errors"
	"testing"
)

// Test success case for FindNode.
func TestFindNodeRequestMessageSuccess(t *testing.T) {
	rpc1 := rpcmarshal.RPC{
		Cmd: "FINO",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "THIS IS ADDRESS"},
		ReqID: newValidRequestID(),
		ID:    *id.NewRandomKademliaID(),
	}
	message := findNodeRequestMessage(rpc1.Contact, rpc1.ReqID, rpc1.ID)
	var rpc2 rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &rpc2)
	if !rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}

// Test fail case for FindNode.
func TestFindNodeRequestMessageFail(t *testing.T) {
	rpc1 := rpcmarshal.RPC{
		Cmd: "FINO",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "THIS IS ADDRESS"},
		ReqID: newValidRequestID(),
	}
	message := findNodeRequestMessage(rpc1.Contact, rpc1.ReqID, *id.NewRandomKademliaID())
	var rpc2 rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &rpc2)
	if rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}

// Test FindNode return success case.
func TestFindNodeRequestReturnSuccess(t *testing.T) {
	var message []byte
	rpcmarshal.RpcMarshal(rpcmarshal.RPC{Content: stored.Value{Data: "something"}}, &message)
	_, err := findNodeRequestReturn(message, nil)
	if isError(err) {
		t.FailNow()
	}
}

// Test FindNode return fail case.
func TestFindNodeRequestReturnTimeout(t *testing.T) {
	_, err := findNodeRequestReturn([]byte{}, errors.New("timeout"))
	if !isError(err) {
		t.FailNow()
	}
}

// Test Find Node response success case.
func TestFindNodeResponseMessageSuccess(t *testing.T) {
	me := contact.GetInstance()
	kademliaID := *id.NewRandomKademliaID()
	bucket.GetInstance().AddContact(contact.Contact{ID: id.NewRandomKademliaID()})
	bucket.GetInstance().AddContact(contact.Contact{ID: id.NewRandomKademliaID()})
	bucket.GetInstance().AddContact(contact.Contact{ID: id.NewRandomKademliaID()})
	bucket.GetInstance().AddContact(contact.Contact{ID: id.NewRandomKademliaID()})
	rpc1 := rpcmarshal.RPC{
		Cmd: "RESP",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "THIS IS ADDRESS"},
		ReqID:  newValidRequestID(),
		KNodes: bucket.GetInstance().FindClosestContacts(&kademliaID, 20),
	}
	message := findNodeResponseMessage(rpc1.Contact, rpc1.ReqID, kademliaID)
	var rpc2 rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &rpc2)
	if !rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}

// Test Find Node response fail case.
func TestFindNodeResponseMessageFail(t *testing.T) {
	kademliaID := *id.NewRandomKademliaID()
	bucket.GetInstance().Me = contact.Contact{ID: id.NewRandomKademliaID()}
	bucket.GetInstance().AddContact(contact.Contact{ID: id.NewRandomKademliaID()})
	bucket.GetInstance().AddContact(contact.Contact{ID: id.NewRandomKademliaID()})
	bucket.GetInstance().AddContact(contact.Contact{ID: id.NewRandomKademliaID()})
	bucket.GetInstance().AddContact(contact.Contact{ID: id.NewRandomKademliaID()})
	rpc1 := rpcmarshal.RPC{
		Cmd: "RESP",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "THIS IS ADDRESS"},
		ReqID: newValidRequestID(),
	}
	message := findNodeResponseMessage(rpc1.Contact, rpc1.ReqID, kademliaID)
	var rpc2 rpcmarshal.RPC
	rpcmarshal.RpcUnmarshal(message, &rpc2)
	if rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}
