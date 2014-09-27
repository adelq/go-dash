package main

import (
	"log"
	"os/exec"
)

func swap() (string, error) {
	swaps, err := exec.Command("cat", "/proc/swaps").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(swaps), nil
}
