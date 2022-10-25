package cli

import (
	"bufio"
	"fmt"
	"os"
)

func CliListener() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		fmt.Println(Handler(scanner.Text()))
	}
}
