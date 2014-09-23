package main

import (
	"log"
	"os/exec"
)

func time() (string, error) {
	time, err := exec.Command("/bin/date").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(time), nil
}
