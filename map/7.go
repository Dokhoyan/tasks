package main

import "errors"

func Invert(m map[string]int) (map[int]string, error) {
	res := make(map[int]string)

	for k, v := range m {
		if _, ok := res[v]; ok {
			return nil, errors.New("duplicate key")
		}
		res[v] = k
	}

	return res, nil
}
