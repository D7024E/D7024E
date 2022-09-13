package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Debug bool  // Debug mode
	err   error // error given from enviroment
)

/**
 * Load enviroment variables
 */
func init() {
	_ = godotenv.Load("../.env")
	// err := nil
	Debug, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		Debug = false
	}
}
