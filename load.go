package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type Load struct {
	Load             []string `json:"load"`
	RunningProcesses string   `json:"runningprocesses"`
	TotalProcesses   string   `json:"totalprocesses"`
}

// load returns the current load average on the system
func load() (systemStruct, error) {
	loadRaw, err := exec.Command("cat", "/proc/loadavg").Output()
	if err != nil {
		log.Fatal(err)
	}

	loadString := string(loadRaw)
	loadAvgSplit := strings.Split(loadString, " ")
	processes := strings.Split(loadAvgSplit[3], "/")
	load := &Load{loadAvgSplit[0:3], processes[0], processes[1]}

	return load, nil
}

// processorCount returns the number of cores in the system
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
