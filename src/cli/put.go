package cli

import (
	"D7024E/node/id"
	"D7024E/node/kademlia/algorithms"
	"D7024E/node/stored"
)

func Put(input []string) string {
	if len(input) != 2 {
		return "Invalid amount of inputs"
	}
	id := *id.NewKademliaID(input[1])
	value := stored.Value{
		Data: input[1],
		ID:   id,
	}
	algorithms.NodeStore(value)
	return id.String()
}
