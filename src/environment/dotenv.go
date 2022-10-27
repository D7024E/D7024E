package environment

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Port       int
	BucketSize int
	Alpha      int
	LogConsole bool // Log Console mode
)

// Load enviroment variables
func init() {
	Alpha = 3
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("[DOTENV] - ", err)
		LogConsole = false
		Port = 4001
		BucketSize = 20
		Alpha = 3
		return
	}

	LogConsole, err = strconv.ParseBool(os.Getenv("LOG_CONSOLE"))
	if err != nil {
		LogConsole = false
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 4001
	}

	BucketSize, err = strconv.Atoi(os.Getenv("BUCKET_SIZE"))
	if err != nil {
		BucketSize = 20
	}

	Alpha, err = strconv.Atoi(os.Getenv("ALPHA"))
	if err != nil {
		Alpha = 3
	}
}
