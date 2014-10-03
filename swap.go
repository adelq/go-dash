package main

import (
	"log"
	"os/exec"
)

func swap() (string, error) {
	swap := exec.Command("cat", "/proc/swaps")
	awk := exec.Command("awk", "{print $1,$2,$3,$4,$5}")

	out, err := swap.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	swap.Start()
	awk.Stdin = out

	swap_out, err := awk.Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(swap_out), nil
}
