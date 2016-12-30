package main

import (
	"log"
	"os/exec"
)

func mem() (string, error) {
	memory := exec.Command("/usr/bin/free", "-tmo")
	awk := exec.Command("awk", `BEGIN {OFS=","} {print $1,$2,$3-$6-$7,$4+$6+$7}`)

	out, err := memory.StdoutPipe()
	if err != nil {
		log.Print(err)
		return "", err
	}
	memory.Start()
	awk.Stdin = out

	memOut, err := awk.Output()
	if err != nil {
		log.Print(err)
		return "", err
	}
	return string(memOut), nil
}
