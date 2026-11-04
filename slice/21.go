package main

import "errors"

func InsertMany(s []int, i int, values []int) ([]int, error) {
	if len(s) < i || i < 0 {
		return []int{}, errors.New("some error")
	}
	if len(values) == 0 {
		return s, nil
	}

	res := make([]int, 0, len(s)+len(values))
	res = append(res, s[:i]...)
	res = append(res, values...)
	res = append(res, s[i:]...)

	return res, nil
}
