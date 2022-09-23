package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Port       int
	Alpha      int
	LogConsole bool  // Log Console mode
	err        error // error given from enviroment
)

/**
 * Load enviroment variables
 */
func init() {
	Port = 4001
	Alpha = 3
	_ = godotenv.Load("../.env")
	// err := nil
	LogConsole, err = strconv.ParseBool(os.Getenv("LOG_CONSOLE"))
	if err != nil {
		LogConsole = false
	}
}
