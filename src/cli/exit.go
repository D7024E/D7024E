package cli

import "os"

func Exit(inputs []string) string {
	if len(inputs) != 1 {
		return "Invalid amount of inputs"
	}
	os.Exit(1)
	return ""
}
