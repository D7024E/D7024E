package cli

import (
	"D7024E/log"
	"bufio"
	"os"
)

func CliListener() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		log.INFO("%v", (scanner.Text()))
	}
}
