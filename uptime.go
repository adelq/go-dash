package main

import (
	"log"
	"os/exec"
)

func uptime() (string, error) {
	uptime, err := exec.Command("uptime").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(uptime), nil
}
