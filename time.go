package main

import (
	"log"
	"os/exec"
)

// time reports the current system date and time
func time() (string, error) {
	time, err := exec.Command("/bin/date").Output()
	if err != nil {
		log.Print(err)
		return "", err
	}
	return string(time), nil
}
