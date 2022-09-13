package log

// https://www.youtube.com/watch?v=p45_9nOpD4k

import (
	"io"
	"log"
	"os"
)

var (
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)

	infoLogger = log.New(multiWriter, "[INFO]    ", log.Ldate|log.Ltime)
	warningLogger = log.New(multiWriter, "[WARNING] ", log.Ldate|log.Ltime)
	errorLogger = log.New(multiWriter, "[ERROR]   ", log.Ldate|log.Ltime|log.Lshortfile)
}

func INFO(message string, v ...any) {
	infoLogger.Printf(message, v...)
}

func WARN(message string, v ...any) {
	warningLogger.Printf(message, v...)
}

func ERROR(message string, v ...any) {
	errorLogger.Printf(message, v...)
}

func FATAL(message string, v ...any) {
	errorLogger.Fatalf(message, v...)
}
