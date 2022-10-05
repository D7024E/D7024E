package cli

import (
	"D7024E/node/kademlia/algorithms"
	"strings"
)

func Handle(text string) string {
	split := strings.Split(text, " ")
	var result string
	switch split[0] {
	case "put":
		result = Put(split)
	case "get":
		result = Get(split[1], algorithms.NodeValueLookup)
	case "exit":
		result = Exit(split)
	default:
		result = "invalid command"
	}
	return result
}
