package main

import (
	"log"
	"os/exec"
	"strings"
)

type Swap struct {
	Filename string `json:"filename"`
	Type     string `json:"type"`
	Size     int    `json:"size"`
	Used     int    `json:"used"`
	Priority int    `json:"priority"`
}

// swap measures swap space and its utilization on the system
func swap() (systemStruct, error) {
	swap := exec.Command("cat", "/proc/swaps")
	awk := exec.Command("awk", "{print $1,$2,$3,$4,$5}")

	out, err := swap.StdoutPipe()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	swap.Start()
	awk.Stdin = out

	swapRaw, err := awk.Output()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	swapOut, err := parseSwap(string(swapRaw))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return swapOut, nil
}

func parseSwap(swap string) ([]Swap, error) {
	swapLines := strings.Split(swap, "\n")
	// Remove header line and trailing newline
	swapLinesTrimmed := swapLines[1 : len(swapLines)-1]
	var result []Swap
	for _, line := range swapLinesTrimmed {
		swap, err := parseSwapLine(line)
		if err != nil {
			return nil, err
		}

		result = append(result, *swap)
	}

	return result, nil
}

func parseSwapLine(line string) (*Swap, error) {
	swapArr := strings.Fields(line)
	intComponents, err := sliceAtoi(swapArr[2:])
	if err != nil {
		return nil, err
	}
	result := &Swap{swapArr[0], swapArr[1], intComponents[0], intComponents[1], intComponents[2]}
	return result, nil
}
