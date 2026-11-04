package main

import "errors"

func ReverseRange(s []int, left, right int) error {
	if left > right || left < 0 || right < 0 || right >= len(s) {
		return errors.New("out of range")
	}

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return nil
}
