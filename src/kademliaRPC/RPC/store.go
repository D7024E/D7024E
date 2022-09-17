package rpc

import (
	"D7024E/node/kademlia"
	"D7024E/node/stored"
)

// Stores values within node instance of values.
func Store(value stored.Value) {
	kademlia.GetInstance().Values.Store([]stored.Value{value})
}
