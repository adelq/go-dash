package main

import (
	"fmt"
	"log"
	"os/exec"
)

func time() string {
	time, err := exec.Command("/bin/date").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(time)
}

func main() {
	time := time()
	fmt.Printf("The date is %s", time)
}
