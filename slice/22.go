package main

import "errors"

func DeleteRange(s []int, left, right int) ([]int, error) {
	if left < 0 || left > right || right > len(s) {
		return s, errors.New("invalid range")
	}

	if left == right {
		return s, nil
	}

	copy(s[left:], s[right:])
	return s[:len(s)-(right-left)], nil
}
