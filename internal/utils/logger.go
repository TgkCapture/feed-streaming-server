package utils

import (
    "log"
    "os"
	"io"
)

var (
    InfoLogger  *log.Logger
    ErrorLogger *log.Logger
)

func InitLogger() {
    logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }

    InfoLogger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

    // To also log to console
    infoMulti := io.MultiWriter(os.Stdout, logFile)
    errorMulti := io.MultiWriter(os.Stderr, logFile)
    InfoLogger = log.New(infoMulti, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(errorMulti, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
