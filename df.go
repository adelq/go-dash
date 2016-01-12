package main

import (
	"log"
	"os/exec"
	"strings"
)

type Df struct {
	Filesystems []Usage `json:"filesystems"`
}

type Usage struct {
	Filesystem  string `json:"filesystem"`
	Size        string `json:"size"`
	Used        string `json:"used"`
	Available   string `json:"available"`
	UsedPercent string `json:"used"`
	Mountpoint  string `json:"mountpoint"`
}

// df reports file system disk space usage
func df() (systemStruct, error) {
	dfRaw, err := exec.Command("df", "-Ph").Output()
	if err != nil {
		log.Fatal(err)
	}

	dfString := string(dfRaw)
	dfSplit := strings.Split(dfString, "\n")
	dfLines := dfSplit[1 : len(dfSplit)-1]

	var result Df
	for _, line := range dfLines {
		usage, err := parseDfLine(line)
		if err != nil {
			log.Fatal(err)
		}

		result.Filesystems = append(result.Filesystems, *usage)
	}

	return result, nil
}

func parseDfLine(line string) (*Usage, error) {
	dfArr := strings.Fields(line)
	result := &Usage{dfArr[0], dfArr[1], dfArr[2], dfArr[3], dfArr[4], dfArr[5]}
	return result, nil
}
