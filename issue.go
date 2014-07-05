package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func issue() (string, string) {
	distro_raw, err := exec.Command("/usr/bin/lsb_release", "-ds").Output()
	if err != nil {
		log.Fatal(err)
	}
	distro := strings.TrimSpace(string(distro_raw))

	kernel_raw, err := exec.Command("/usr/bin/uname", "-r").Output()
	if err != nil {
		log.Fatal(err)
	}
	kernel := strings.TrimSpace(string(kernel_raw))

	return distro, kernel
}

func main() {
	distro, kernel := issue()
	fmt.Printf("The distro is %s, and the kernel is %s\n", distro, kernel)
}
