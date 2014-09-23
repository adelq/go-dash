package main

import (
	"log"
	"os"
)

func hostname() (string, error) {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return name, nil
}
