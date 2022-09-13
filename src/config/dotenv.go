package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Debug bool

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Debug, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatal("Error converting to Debug to bool")
	}
}
