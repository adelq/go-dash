package main

import (
	"log"
	"os/exec"
	"strings"
)

type Issue struct {
	Distro string `json:"distro"`
	Kernel string `json:"kernel"`
}

// issue returns the Linux distribution name and kernel version in use
func issue() (systemStruct, error) {
	distroRaw, err := exec.Command("lsb_release", "-ds").Output()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	distro := strings.TrimSpace(string(distroRaw))

	kernelRaw, err := exec.Command("uname", "-r").Output()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	kernel := strings.TrimSpace(string(kernelRaw))

	return &Issue{distro, kernel}, nil
}
