package log

// https://www.youtube.com/watch?v=p45_9nOpD4k

import (
	"io"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("log/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)

	InfoLogger = log.New(multiWriter, "[INFO]    ", log.Ldate|log.Ltime)
	WarningLogger = log.New(multiWriter, "[WARNING] ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(multiWriter, "[ERROR]   ", log.Ldate|log.Ltime|log.Lshortfile)
}

func INFO(message string, v ...any) {
	InfoLogger.Printf(message, v...)
}

func WARN(message string, v ...any) {
	WarningLogger.Printf(message, v...)
}

func ERROR(message string, v ...any) {
	ErrorLogger.Printf(message, v...)
}
