package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type CPUInfo struct {
	Architecture string  `json:"architecture"`
	CPUs         int64   `json:"cpus"`
	Sockets      int64   `json:"sockets"`
	CPUFamily    int64   `json:"cpu_family"`
	BogoMIPS     float64 `json:"bogomips"`
}

func cpuinfo() (systemStruct, error) {
	cpuInfoRaw, err := exec.Command("lscpu").Output()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	cpuInfoSplit := strings.Split(string(cpuInfoRaw), "\n")

	result := &CPUInfo{}
	result.Architecture = strings.Fields(cpuInfoSplit[0])[1]
	result.CPUs, _ = strconv.ParseInt(strings.Fields(cpuInfoSplit[3])[1], 0, 64)
	result.Sockets, _ = strconv.ParseInt(strings.Fields(cpuInfoSplit[7])[1], 0, 64)
	result.CPUFamily, _ = strconv.ParseInt(strings.Fields(cpuInfoSplit[10])[2], 0, 64)
	result.BogoMIPS, _ = strconv.ParseFloat(strings.Fields(cpuInfoSplit[14])[1], 64)
	return result, nil
}
