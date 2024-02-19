package logging

import (
	"log"
	"os"
)

func Setup(logPath string) {
	// Create a file for logging
	file, err := os.OpenFile(logPath+"/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	// Set log output to file
	log.SetOutput(file)
}
