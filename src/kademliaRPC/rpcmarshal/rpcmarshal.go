package rpcmarshal

import (
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"encoding/json"
)

type RPC struct {
	Cmd     string
	Contact contact.Contact
	ReqID   string
	ID      id.KademliaID
	Content stored.Value
}

// A basic test is bellow, move it to main for testing.
// Takes a rpc struct and marshalls it to JSON, writes it in res.
func RpcMarshal(rpc RPC, res *[]byte) {
	rpcJSON, _ := json.Marshal(rpc)
	*res = rpcJSON
}

// Takes a marshaled message and unmarshals it to a RPC struct and writes it to res.
func RpcUnmarshal(msg []byte, res *RPC) {
	json.Unmarshal(msg, res)
}

// var testJSON []byte
// var testOut network.RPC
// testid := Contact{
// 	ID:       node.NewRandomKademliaID(),
// 	Address:  "THIS IS ADDRESS",
// 	Distance: node.NewRandomKademliaID(),
// }
// test := network.RPC{
// 	Cmd:     "test",
// 	Contact: []string{"1", "2", "3"},
// 	Id:      testid,
// }
// network.RpcMarshal(test, &testJSON)
// fmt.Println(testJSON)
// network.RpcUnmarshal(testJSON, &testOut)
// fmt.Println(testOut)
