package main

import (
	"log"
	"os/exec"
)

func arp() (string, error) {
	arp, err := exec.Command("arp").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(arp), nil
}
