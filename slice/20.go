package main

import "fmt"

func InsertAt(s []int, i, value int) ([]int, error) {
	if i < 0 || i > len(s) {
		return s, fmt.Errorf("index out of range")
	}
	s = append(s, 0)

	copy(s[i+1:], s[i:])
	s[i] = value
	return s, nil
}
