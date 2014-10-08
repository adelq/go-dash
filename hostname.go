package main

import (
	"log"
	"os"
)

// hostname returns the system's current host name
func hostname() (string, error) {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return name, nil
}
