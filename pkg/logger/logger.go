package logger

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(message string, keysAndValues ...interface{}) {
	infoLogger.Println(append([]interface{}{message}, keysAndValues...)...)
}

func Error(message string, keysAndValues ...interface{}) {
	errorLogger.Println(append([]interface{}{message}, keysAndValues...)...)
}
