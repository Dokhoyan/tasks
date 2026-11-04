package main

import "unicode"

func TrimUnicodeSpace(s string) string {
	runs := []rune(s)
	left := 0
	right := len(runs) - 1

	for left <= right && unicode.IsSpace(runs[left]) {
		left++
	}

	for left <= right && unicode.IsSpace(runs[right]) {
		right--
	}

	return string(runs[left : right+1])
}
