package main

import (
	"fmt"
	log "github.com/charmbracelet/log"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Read the environment variable
	listStr := os.Getenv("THIS_IS_FINE_STDOUT_LOGGING")
	// Check if the environment variable is not empty
	if listStr != "" {
		logItems := strings.Split(listStr, ",")
		for _, logItem := range logItems {
			logMessages := strings.Split(logItem, ":")
			if len(logMessages) != 2 {
				log.Fatal("Please define log level via 'INFO:example-log-message'. Supported levels are: INFO, WARN, ERROR and DEBUG")
			}
			switch sev := logMessages[0]; sev {
			case "INFO":
				log.Info(logMessages[1])
			case "WARN":
				log.Warn(logMessages[1])
			case "ERROR":
				log.Error(logMessages[1])
			case "DEBUG":
				log.Debug(logMessages[1])
			default:
				log.Print(logMessages[1])
			}
		}
	} else {
		fmt.Println("Environment variable THIS_IS_FINE_STDOUT_LOGGING is not set or is empty. Allowed format: THIS_IS_FINE_STDOUT_LOGGING=<LEVEL>:<MESSAGE>,<LEVEL>:<MESSAGE>")
	}

}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
