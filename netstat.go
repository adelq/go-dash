package main

import (
	"log"
	"os/exec"
)

// netstat prints all current network connections and routing tables
func netstat() (string, error) {
	netstat, err := exec.Command("netstat", "-ntu").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(netstat), nil
}
