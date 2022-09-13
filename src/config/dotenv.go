package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Debug bool

func init() {
	err := godotenv.Load("../.env")

	Debug, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		Debug = false
	}
}
