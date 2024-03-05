package main

import (
	"fmt"
	log "github.com/charmbracelet/log"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Check if the environment variable is not empty
	// Read the environment variable
	listStr := os.Getenv("THIS_IS_FINE_STDOUT_LOGGING")
	if listStr != "" {
		// Split the string based on the comma delimiter to get a slice of strings
		list := strings.Split(listStr, ",")

		// Iterate through the list and print each element
		for _, item := range list {
			log.Info(item)
		}
	} else {
		fmt.Println("Environment variable THIS_IS_FINE_STDOUT_LOGGING is not set or is empty.")
	}

}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
