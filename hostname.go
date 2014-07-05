package main

import (
	"fmt"
	"log"
	"os"
)

func hostname() string {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return name
}

func main() {
	fmt.Println(hostname())
}
