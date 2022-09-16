package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	LogConsole bool  // Debug mode
	err        error // error given from enviroment
)

/**
 * Load enviroment variables
 */
func init() {
	_ = godotenv.Load("../.env")
	// err := nil
	LogConsole, err = strconv.ParseBool(os.Getenv("LOG_CONSOLE"))
	if err != nil {
		LogConsole = false
	}
}
