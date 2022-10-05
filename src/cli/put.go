package cli

import (
	"D7024E/node/id"
	"D7024E/node/stored"
)

type NodeStore func(stored.Value) bool

// Returns a hash of the input string if it was stored successfully, otherwise returns "".
func Put(input string, NS NodeStore) string {
	id := *id.NewKademliaID(input)
	value := stored.Value{
		Data: input,
		ID:   id,
	}
	var res bool = NS(value)
	if res {
		return id.String()
	} else {
		return ""
	}
}
