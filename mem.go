package main

import (
	"log"
	"os/exec"
)

func mem() (string, error) {
	mem, err := exec.Command("/usr/bin/free", "-tmo").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(mem), nil
}
