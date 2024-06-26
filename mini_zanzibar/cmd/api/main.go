package main

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"miniZanzibar/internal/server"
	"os"
)

var logFile *os.File

func setupLogFile() error {
	fmt.Println("Setting up log file")

	// Open a file for writing (create it if not exist, append to it if it does)
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}

	logFile = file
	return nil
}

func init() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up logrus
	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}

	// Create or open the log file
	err = setupLogFile()
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}

	// Set logrus output to the file
	log.SetOutput(logFile)
	log.SetLevel(logLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	// Defer closing the log file to ensure it gets closed when the main function exits
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error("Error closing log file: ", err)
		}
	}(logFile)

	server := server.NewServer()
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
