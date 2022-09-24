package cli

import (
	"strings"
)

func Handle(text string) string {
	split := strings.Split(text, " ")
	var result string
	switch split[0] {
	case "put":
		result = Put(split)
	case "get":
		result = Get(split)
	case "exit":
		result = Exit(split)
	default:
		result = "Invalid command"
	}
	return result
}
