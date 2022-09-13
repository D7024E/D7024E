package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Debug bool // Debug mode

/**
 * Load enviroment variables
 */
func init() {
	err := godotenv.Load("../.env")

	Debug, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		Debug = false
	}
}
