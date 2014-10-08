package main

import (
	"log"
	"os/exec"
)

// ping reports the average ping time to a given url
func pingTime(url string) (string, error) {
	pingOutput, err := exec.Command("ping", "-qc1", url).Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(pingOutput), nil
}
