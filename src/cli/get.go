package cli

import (
	"D7024E/node/id"
	"D7024E/node/stored"
)

type NodeValueLookup func(id.KademliaID) (stored.Value, error)

func Get(input string, NVL NodeValueLookup) string {
	valueID, err := id.String2KademliaID(input)
	if err != nil {
		return err.Error()
	}

	value, err := NVL(*valueID)
	if err != nil {
		return err.Error()
	}

	return value.Data
}
