package main

import (
	"log"
	"os/exec"
	"strings"
)

type Netstat struct {
	Protocol       string `json:"protocol"`
	LocalAddress   string `json:"local_address"`
	ForeignAddress string `json:"foreign_address"`
	State          string `json:"state"`
}

// netstat prints all current network connections and routing tables
func netstat() (systemStruct, error) {
	netstat, err := exec.Command("netstat", "-ntu").Output()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	netstatOut, err := parseNetstat(string(netstat))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return netstatOut, nil
}

func parseNetstat(netstat string) ([]Netstat, error) {
	netstatLines := strings.Split(netstat, "\n")
	// Remove 2 header lines and trailing newline
	nsLinesTrimmed := netstatLines[2 : len(netstatLines)-1]
	var result []Netstat
	for _, line := range nsLinesTrimmed {
		ns, err := parseNetstatLine(line)
		if err != nil {
			return nil, err
		}

		result = append(result, *ns)
	}

	return result, nil
}

func parseNetstatLine(line string) (*Netstat, error) {
	nsArr := strings.Fields(line)
	// TODO: Interpret Recv-Q and Send-Q columns
	result := &Netstat{nsArr[0], nsArr[3], nsArr[4], nsArr[5]}
	return result, nil
}
