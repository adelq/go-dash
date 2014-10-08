package main

import (
	"log"
	"os/exec"
)

// df reports file system disk space usage
func df() (string, error) {
	df, err := exec.Command("df", "-Ph").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(df), nil
}
