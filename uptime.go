package main

import (
	"log"
	"os/exec"
)

// uptime reports how long the system has been running
func uptime() (string, error) {
	uptime, err := exec.Command("uptime").Output()
	if err != nil {
		log.Print(err)
		return "", err
	}

	return string(uptime), nil
}
