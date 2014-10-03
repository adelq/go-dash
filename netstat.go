package main

import (
	"log"
	"os/exec"
)

func netstat() (string, error) {
	netstat, err := exec.Command("netstat", "-ntu").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(netstat), nil
}
