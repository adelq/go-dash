package main

import (
	"log"
	"os/exec"
)

func df() (string, error) {
	df, err := exec.Command("df", "-Ph").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(df), nil
}
