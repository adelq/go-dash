package main

import (
	"log"
	"os/exec"
	"strings"
)

// issue returns the Linux distribution name and kernel version in use
func issue() (string, string, error) {
	distro_raw, err := exec.Command("lsb_release", "-ds").Output()
	if err != nil {
		log.Fatal(err)
	}
	distro := strings.TrimSpace(string(distro_raw))

	kernel_raw, err := exec.Command("uname", "-r").Output()
	if err != nil {
		log.Fatal(err)
	}
	kernel := strings.TrimSpace(string(kernel_raw))

	return distro, kernel, nil
}
