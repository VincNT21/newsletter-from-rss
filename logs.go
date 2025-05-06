package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Create different loggers for different purposes
var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
)

// Create a "had error" check and an error list
var hadError bool
var errorComponents []string

func setupLogging() {
	// Default to stderr if file creation fails
	logWriter := os.Stderr

	// Try to create log file
	logPath := getFilePath("app.log")

	// Make sure the directory exists
	dir := filepath.Dir(logPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("Failed to create log directory: %v\n", err)
	}

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err == nil {
		logWriter = file
		// We don't close the file here because we'll use it thoughout the program
		// The OS will close it when the program exits
	} else {
		fmt.Printf("failed to open log file: %v\n", err)
	}

	// Create loggers with prefixes
	infoLogger = log.New(logWriter, "INFO: ", log.Ldate|log.Ltime)
	errorLogger = log.New(logWriter, "ERROR: ", log.Ldate|log.Ltime)
	debugLogger = log.New(logWriter, "DEBUG: ", log.Ldate|log.Ltime)
}

func logError(errMessage, component string) {
	errorLogger.Println(errMessage)
	hadError = true
	errorComponents = append(errorComponents, component)
}
