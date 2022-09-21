package rpcmarshal

import (
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"encoding/json"
	"reflect"
)

type RPC struct {
	Cmd     string
	Contact contact.Contact
	ReqID   string
	ID      id.KademliaID
	Content stored.Value
	KNodes  []contact.Contact
}

func (r1 *RPC) Equals(r2 *RPC) bool {
	var res bool
	if !(reflect.ValueOf(*r1).FieldByName("Cmd").IsZero()) {
		res = r1.Cmd == r2.Cmd
	}
	if !(reflect.ValueOf(*r1).FieldByName("Contact").IsZero()) {
		res = res && r1.Contact.Equals(&r2.Contact)
	}
	if !(reflect.ValueOf(*r1).FieldByName("ReqID").IsZero()) {
		res = res && (r2.ReqID == r1.ReqID)
	}
	if !(reflect.ValueOf(*r1).FieldByName("ID").IsZero()) {
		res = res && r1.ID.Equals(&r2.ID)
	}
	if !(reflect.ValueOf(*r1).FieldByName("Content").IsZero()) {
		res = res && r1.Content.Equals(&r2.Content)
	}
	return res
}

// A basic test is bellow, move it to main for testing.
// Takes a rpc struct and marshalls it to JSON, writes it in res.
func RpcMarshal(rpc RPC, res *[]byte) error {
	rpcJSON, err := json.Marshal(rpc)
	*res = rpcJSON
	return err
}

// Takes a marshaled message and unmarshals it to a RPC struct and writes it to res.
func RpcUnmarshal(msg []byte, res *RPC) {
	json.Unmarshal(msg, res)
}
