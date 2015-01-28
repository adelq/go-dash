package main

import (
	"log"
	"os/exec"
)

// swap measures swap space and its utilization on the system
func swap() (string, error) {
	swap := exec.Command("cat", "/proc/swaps")
	awk := exec.Command("awk", "{print $1,$2,$3,$4,$5}")

	out, err := swap.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	swap.Start()
	awk.Stdin = out

	swapOut, err := awk.Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(swapOut), nil
}
