package rpcmarshal

import (
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"testing"
	"time"
)

// Test marshal and unmarshal rpc.
func TestRpcMarshalSuccess(t *testing.T) {
	rpc := RPC{
		Cmd: "THIS IS CMD",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "0.0.0.0",
		},
		ID: *id.NewRandomKademliaID(),
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

// Test fail case for marshaling.
func TestRpcMarshalFail(t *testing.T) {
	rpc := RPC{
		Cmd: "THIS IS CMD 2",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "0.0.0.0",
		},
		ID: *id.NewRandomKademliaID(),
		Content: stored.Value{
			Data:   "THIS IS DATA 2",
			ID:     *id.NewRandomKademliaID(),
			Ttl:    time.Second,
			DeadAt: time.Now().Add(time.Second),
		},
	}
	message := []byte(`{
		"Cmd":"THIS IS CMD",
		"Contact":{
			"ID":[33,15,199,187,129,134,57,172,72,164,198,175,162,241,88,26,139,149,37,226],
			"Address":"0.0.0.0"},
		"ID":[15,218,104,146,127,43,47,248,54,247,53,120,219,15,165,76,41,247,253,146],
		"Content":{
			"name":"THIS IS DATA",
			"id":[141,146,202,67,241,147,222,228,127,89,21,73,245,151,168,17,200,250,103,171],
			"ttl":1000000000,"deadAt":"2022-09-21T15:31:33.7349182+02:00"},
			"KNodes":null}`)
	var rpc2 RPC
	RpcUnmarshal(message, &rpc2)

	if rpc.Equals(&rpc2) {
		t.FailNow()
	}
}

// Test RPC equals for success case.
func TestRPCEqualsSuccess(t *testing.T) {
	rpc := RPC{
		Cmd: "THIS IS CMD",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "0.0.0.0",
		},
		ID: *id.NewRandomKademliaID(),
		Content: stored.Value{
			Data:   "THIS IS DATA",
			ID:     *id.NewRandomKademliaID(),
			Ttl:    time.Second,
			DeadAt: time.Now().Add(time.Second),
		},
	}
	if !rpc.Equals(&rpc) {
		t.FailNow()
	}
}

// Test RPC Equals with 2 different RPC.
func TestRPCEqualsFail(t *testing.T) {
	rpc := RPC{
		Cmd: "THIS IS CMD",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "0.0.0.0",
		},
		ID: *id.NewRandomKademliaID(),
		Content: stored.Value{
			Data:   "THIS IS DATA",
			ID:     *id.NewRandomKademliaID(),
			Ttl:    time.Second,
			DeadAt: time.Now().Add(time.Second),
		},
	}
	rpc2 := RPC{
		Cmd: "THIS IS CMD 2",
		Contact: contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: "0.0.0.0",
		},
		ID: *id.NewRandomKademliaID(),
		Content: stored.Value{
			Data:   "THIS IS DATA 2",
			ID:     *id.NewRandomKademliaID(),
			Ttl:    time.Second,
			DeadAt: time.Now().Add(time.Second),
		},
	}
	if rpc.Equals(&rpc2) {
		t.FailNow()
	} else if rpc2.Equals(&rpc) {
		t.FailNow()
	}
}

// Test RPC Equals for fail case with missing Cmd.
func TestRPCEqualsFailMissingCmd(t *testing.T) {
	rpc := RPC{}
	rpc2 := rpc
	rpc2.Cmd = "THIS IS CMD"
	if rpc.Equals(&rpc2) {
		t.FailNow()
	} else if rpc2.Equals(&rpc) {
		t.FailNow()
	}
}

// Test RPC Equals for fail case with missing Contact.
func TestRPCEqualsFailMissingContact(t *testing.T) {
	rpc := RPC{}
	rpc2 := rpc
	rpc2.Contact = contact.Contact{
		ID:      id.NewRandomKademliaID(),
		Address: "0.0.0.0",
	}
	if rpc.Equals(&rpc2) {
		t.FailNow()
	} else if rpc2.Equals(&rpc) {
		t.FailNow()
	}
}

// Test RPC Equals for fail case with missing ID.
func TestRPCEqualsFailMissingID(t *testing.T) {
	rpc := RPC{}
	rpc2 := rpc
	rpc2.ID = *id.NewRandomKademliaID()
	if rpc.Equals(&rpc2) {
		t.FailNow()
	} else if rpc2.Equals(&rpc) {
		t.FailNow()
	}
}

// Test RPC Equals for fail case with missing Content.
func TestRPCEqualsFailMissingContent(t *testing.T) {
	rpc := RPC{}
	rpc2 := rpc
	rpc2.Content = stored.Value{
		Data:   "THIS IS DATA",
		ID:     *id.NewRandomKademliaID(),
		Ttl:    time.Second,
		DeadAt: time.Now().Add(time.Second),
	}
	if rpc.Equals(&rpc2) {
		t.FailNow()
	} else if rpc2.Equals(&rpc) {
		t.FailNow()
	}
}

// Test RPC Equals for fail case with missing Content.
func TestRPCEqualsFailMissingKNodes(t *testing.T) {
	rpc := RPC{}
	rpc2 := rpc
	rpc2.KNodes = []contact.Contact{{}}
	if rpc.Equals(&rpc2) {
		t.FailNow()
	} else if rpc2.Equals(&rpc) {
		t.FailNow()
	}
}
