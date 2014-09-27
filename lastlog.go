package main

import (
	"log"
	"os/exec"
)

func lastlog() (string, error) {
	users, err := exec.Command("/usr/bin/lastlog", "--time", "365").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(users), nil
}
