package main

import "errors"

func RemoveAtFast(s []int, i int) ([]int, error) {
	if i < 0 || i >= len(s) {
		return nil, errors.New("index out of range")
	}

	s[i] = s[len(s)-1]

	return s[:len(s)-1], nil
}
