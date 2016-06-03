package main

import "strconv"

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func sliceAtof(sa []string) ([]float64, error) {
	si := make([]float64, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.ParseFloat(a, 64)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
