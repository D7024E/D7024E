package cli

import (
	"D7024E/node/id"
	"D7024E/node/kademlia/algorithms"
)

func Get(input []string) string {
	if len(input) != 2 {
		return "invalid amount of arguments"
	}

	valueID, err := id.String2KademliaID(input[1])
	if err != nil {
		return err.Error()
	}

	value, err := algorithms.NodeValueLookup(*valueID)
	if err != nil {
		return err.Error()
	}

	return value.Data
}
