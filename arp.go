package main

import (
	"log"
	"os/exec"
	"strings"
)

type Arp struct {
	Address   string `json:"address"`
	HWtype    string `json:"hwtype"`
	HWaddress string `json:"hwaddress"`
	Flags     string `json:"flags"`
	Interface string `json:"iface"`
}

// arp displays the kernel's IPv4 network neighbor cache
func arp() (systemStruct, error) {
	arpRaw, err := exec.Command("arp").Output()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	arpString := string(arpRaw)
	arpLine := strings.Split(arpString, "\n")[1]
	arpArr := strings.Fields(arpLine)
	arp := &Arp{arpArr[0], arpArr[1], arpArr[2], arpArr[3], arpArr[4]}

	return arp, nil
}
