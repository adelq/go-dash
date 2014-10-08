package main

import (
	"log"
	"os/exec"
)

// arp displays the kernel's IPv4 network neighbor cache
func arp() (string, error) {
	arp, err := exec.Command("arp").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(arp), nil
}
