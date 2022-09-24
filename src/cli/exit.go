package cli

import "os"

func Exit(inputs []string) string {
	if len(inputs) != 1 {
		return "invalid amount of arguments"
	}
	os.Exit(1)
	return ""
}
