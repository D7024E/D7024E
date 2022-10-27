package environment

import (
	"D7024E/log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Port       int
	Alpha      int
	LogConsole bool // Log Console mode
)

/**
 * Load enviroment variables
 */
func init() {
	Alpha = 3
	err := godotenv.Load("../.env")
	if err != nil {
		log.ERROR("[DOTENV] - [%v]", err)
		LogConsole = false
		Port = 4001
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
}
