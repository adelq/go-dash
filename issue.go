package main

import (
	"log"
	"os/exec"
	"strings"
)

// issue returns the Linux distribution name and kernel version in use
func issue() (string, string, error) {
	distroRaw, err := exec.Command("lsb_release", "-ds").Output()
	if err != nil {
		log.Fatal(err)
	}
	distro := strings.TrimSpace(string(distroRaw))

	kernelRaw, err := exec.Command("uname", "-r").Output()
	if err != nil {
		log.Fatal(err)
	}
	kernel := strings.TrimSpace(string(kernelRaw))

	return distro, kernel, nil
}
