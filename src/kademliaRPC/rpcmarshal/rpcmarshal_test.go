package rpcmarshal

import (
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"fmt"
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
			Data: "THIS IS DATA",
			TTL:  time.Second,
		},
		Acknowledge: true,
	}
	var message []byte
	var rpc2 RPC
	err := RpcMarshal(rpc, &message)
	if err != nil {
		t.FailNow()
	}
	RpcUnmarshal(message, &rpc2)

	fmt.Println(rpc)
	fmt.Println(rpc2)

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
			TTL:    time.Second,
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
			"KNodes":null},
		"Acknowledge":false`)
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
			TTL:    time.Second,
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
			TTL:    time.Second,
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
			TTL:    time.Second,
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
		TTL:    time.Second,
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

// test equals for kNodes when length of kNodes is 1 and 0
func TestRPCEqualsKNodesNotSameLength(t *testing.T) {
	rpc1 := RPC{}
	rpc2 := RPC{}
	rpc1.KNodes = []contact.Contact{{ID: id.NewRandomKademliaID()}}
	if rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}

// test equals for kNodes when length of kNodes is 2 and 1
func TestRPCEqualsKNodesNotSameLength2Elements(t *testing.T) {
	rpc1 := RPC{}
	rpc2 := RPC{}
	rpc1.KNodes = []contact.Contact{{ID: id.NewRandomKademliaID()}, {ID: id.NewRandomKademliaID()}}
	rpc2.KNodes = []contact.Contact{{ID: id.NewRandomKademliaID()}}
	if rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}

// Test if two different rpc based on KNodes is equal.
func TestRPCEqualsKNodesSuccess(t *testing.T) {
	rpc1 := RPC{}
	rpc2 := RPC{}
	id := id.NewRandomKademliaID()
	rpc1.KNodes = []contact.Contact{{ID: id}}
	rpc2.KNodes = []contact.Contact{{ID: id}}
	if !rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}

// Test if the same RPC is equal based on KNodes
func TestRPCEqualsKNodesSuccessSameRPC(t *testing.T) {
	rpc1 := RPC{}
	id := id.NewRandomKademliaID()
	rpc1.KNodes = []contact.Contact{{ID: id}}
	if !rpc1.Equals(&rpc1) {
		t.FailNow()
	}
}

// Test if two different KNodes ids are different
func TestRPCEqualsKNodesFail(t *testing.T) {
	rpc1 := RPC{}
	rpc2 := RPC{}
	rpc1.KNodes = []contact.Contact{{ID: id.NewRandomKademliaID()}}
	rpc2.KNodes = []contact.Contact{{ID: id.NewRandomKademliaID()}}
	if rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}

// Test KNodes on one rpc with one element and one empty
func TestRPCEqualsKNodesFailOneEmpty(t *testing.T) {
	rpc1 := RPC{}
	rpc2 := RPC{}
	rpc1.KNodes = []contact.Contact{{ID: id.NewRandomKademliaID()}}
	rpc2.KNodes = []contact.Contact{{}}
	if rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}

func TestRPCEqualsKNodesSuccessTwoElementsInEach(t *testing.T) {
	rpc1 := RPC{}
	rpc2 := RPC{}
	id1 := id.NewRandomKademliaID()
	id2 := id.NewRandomKademliaID()
	rpc1.KNodes = []contact.Contact{{ID: id1}, {ID: id2}}
	rpc2.KNodes = []contact.Contact{{ID: id1}, {ID: id2}}

	if !rpc1.Equals(&rpc2) {
		t.FailNow()
	}
}
