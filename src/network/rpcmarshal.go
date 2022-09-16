package network

import (
	"D7024E/node"
	"encoding/json"
)

type RPC struct {
	cmd     string
	contact []string
	id      *node.KademliaID
}

func Marshall(res *[]byte, rpc RPC) {
	rpcJSON, _ := json.Marshal(rpc)
	*res = rpcJSON
}
