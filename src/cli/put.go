package cli

import (
	"D7024E/node/id"
	"D7024E/node/stored"
	"time"
)

type NodeStore func(stored.Value) bool
type RefreshAlgorithm func(stored.Value)

// Returns a hash of the input string if it was stored successfully, otherwise returns "".
func Put(input string, NS NodeStore, RA RefreshAlgorithm) string {
	id := *id.NewKademliaID(input)
	value := stored.Value{
		Data: input,
		ID:   id,
		TTL:  time.Minute,
	}
	res := NS(value)
	if res {
		go RA(value)
		return id.String()
	} else {
		return ""
	}
}
