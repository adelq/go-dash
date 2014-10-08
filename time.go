package main

import (
	"log"
	"os/exec"
)

// time reports the current system date and time
func time() (string, error) {
	time, err := exec.Command("/bin/date").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(time), nil
}
