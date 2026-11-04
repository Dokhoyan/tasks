package main

import "errors"

func CloneRange(s []int, left, right int) ([]int, error) {
	if left > right || left < 0 || right > len(s) {
		return nil, errors.New("invalid range")
	}

	res := make([]int, right-left)
	copy(res, s[left:right])
	return res, nil
}
