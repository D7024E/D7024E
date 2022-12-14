package rpcmarshal

import (
	"D7024E/log"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"bytes"
	"encoding/json"
	"reflect"
)

type RPC struct {
	Cmd         string            `json:"cmd"`
	Contact     contact.Contact   `json:"contact"`
	ID          id.KademliaID     `json:"id"`
	Content     stored.Value      `json:"content"`
	KNodes      []contact.Contact `json:"knodes"`
	Acknowledge bool              `json:"acknowledge"`
}

// Check if two RPC are equal, return true if they are otherwise false.
func (r1 *RPC) Equals(r2 *RPC) bool {
	var res bool
	// Cmd
	if reflect.ValueOf(*r1).IsZero() && reflect.ValueOf(*r2).IsZero() {
		return true
	}

	if !(reflect.ValueOf(*r1).FieldByName("Cmd").IsZero()) {
		res = r1.Cmd == r2.Cmd
	} else if !(reflect.ValueOf(*r2).FieldByName("Cmd").IsZero()) {
		return false
	}

	// Contact
	if !(reflect.ValueOf(*r1).FieldByName("Contact").IsZero()) {
		res = res && r1.Contact.Equals(&r2.Contact)
	} else if !(reflect.ValueOf(*r2).FieldByName("Contact").IsZero()) {
		return false
	}

	// ID
	if !(reflect.ValueOf(*r1).FieldByName("ID").IsZero()) {
		res = res && r1.ID.Equals(&r2.ID)
	} else if !(reflect.ValueOf(*r2).FieldByName("ID").IsZero()) {
		return false
	}

	// Content
	if !(reflect.ValueOf(*r1).FieldByName("Content").IsZero()) {
		res = res && r1.Content.Equals(&r2.Content)
	} else if !(reflect.ValueOf(*r2).FieldByName("Content").IsZero()) {
		return false
	}

	// KNodes
	if !(reflect.ValueOf(*r1).FieldByName("KNodes").IsZero()) {
		if len(r1.KNodes) != len(r2.KNodes) {
			return false
		} else {
			for i, c := range r1.KNodes {
				if !c.Equals(&r2.KNodes[i]) {
					return false
				}
			}
			return true
		}
	} else if !(reflect.ValueOf(*r2).FieldByName("KNodes").IsZero()) {
		return false
	}

	// Acknowledge
	if !(reflect.ValueOf(*r1).FieldByName("Acknowledge").IsZero()) {
		res = res && (r1.Acknowledge == r2.Acknowledge)
	} else if !(reflect.ValueOf(*r2).FieldByName("Acknowledge").IsZero()) {
		return false
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
	msg = bytes.Trim(msg, "\x00")
	err := json.Unmarshal(msg, res)
	if err != nil {
		log.ERROR("UNMARSHAL - %v", err)
	}
}
