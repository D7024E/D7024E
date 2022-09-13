package log

// https://www.youtube.com/watch?v=p45_9nOpD4k

import (
	"D7024E/config"
	"io"
	"log"
	"os"
)

var (
	multiWriter   io.Writer   // MultiWrtier for log
	infoLogger    *log.Logger // Logger for info messages
	warningLogger *log.Logger // Logger for warnings messages
	errorLogger   *log.Logger // Logger for error messages including fatal
)

/**
 * Initialize the log writers as multiwriters meaning that they both write to
 * file logs.txt and to the console being the os.Stdout.
 */
func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	if config.Debug {
		multiWriter = io.MultiWriter(os.Stdout, file)
	} else {
		multiWriter = io.MultiWriter(file)
	}

	infoLogger = log.New(multiWriter, "[INFO]    ", log.Ldate|log.Ltime)
	warningLogger = log.New(multiWriter, "[WARNING] ", log.Ldate|log.Ltime)
	errorLogger = log.New(multiWriter, "[ERROR]   ", log.Ldate|log.Ltime|log.Lshortfile)
}

/**
 * Log info messages.
 */
func INFO(message string, v ...any) {
	infoLogger.Printf(message, v...)
}

/**
 * Log warning messages.
 */
func WARN(message string, v ...any) {
	warningLogger.Printf(message, v...)
}

/**
 * Log error messages.
 */
func ERROR(message string, v ...any) {
	errorLogger.Printf(message, v...)
}

/**
 * Log fatal messages which then exit.
 */
func FATAL(message string, v ...any) {
	errorLogger.Fatalf(message, v...)
}
