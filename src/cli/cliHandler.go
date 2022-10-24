package cli

import (
	"D7024E/node/kademlia/algorithms"
	"os"
	"strings"
)

func Handle(text string) string {
	cmd, content := parseInput(text)

	var result string
	switch cmd {
	case "put":
		result = Put(content, algorithms.NodeStore)
	case "get":
		result = Get(content, algorithms.NodeValueLookup)
	case "exit":
		os.Exit(1)
	default:
		result = "invalid command"
	}
	return result
}

func parseInput(input string) (string, string) {
	split := strings.Split(input, " ")

	var cmd string = split[0]
	var content string
	for i := 1; i < len(split); i++ {
		content += split[i] + " "
	}
	content = content[:len(content)-1]
	return cmd, content
}
