package main

import (
	"log"
	"os/exec"
)

// lastlog reports the most recent login of all users on the system
func lastlog() (string, error) {
	users := exec.Command("/usr/bin/lastlog", "--time", "365")
	awk := exec.Command("/usr/bin/awk", `{print $1","$3","$4" "$5" "$6" "$7" "$8}`)

	out, err := users.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	users.Start()
	awk.Stdin = out

	usersOut, err := awk.Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(usersOut), nil
}
