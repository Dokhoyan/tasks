package main

import "errors"

func Splice(s []int, start, deleteCount int, values []int) ([]int, error) {
	if start > len(s) {
		return s, errors.New("invalid start")
	}

	if len(s[start:]) < deleteCount {
		return nil, errors.New("not enough elements")
	}

	oldLen := len(s)
	newLen := oldLen - deleteCount + len(values)
	if oldLen < newLen {
		s = append(s, make([]int, newLen-oldLen)...)
	}

	copy(s[start+len(values):], s[start+deleteCount:])
	copy(s[start:], values)

	return s[:newLen], nil
}
