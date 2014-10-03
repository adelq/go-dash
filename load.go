package main

import (
	"log"
	"os/exec"
	"strconv"
)

func load() (string, error) {
	load, err := exec.Command("cat", "/proc/loadavg").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(load), nil
}

func processorCount() (int64, error) {
	processors, err := exec.Command("grep", "-c", "^processor", "/proc/cpuinfo").Output()
	if err != nil {
		log.Fatal(err)
	}

	numberProcessors, err := strconv.ParseInt(string(processors), 0, 64)
	if err != nil {
		log.Fatal(err)
	}
	return numberProcessors, nil
}
