package main

import "errors"

func Split(s, sep string) ([]string, error) {
	if len(sep) == 0 {
		return []string{}, errors.New("sep is empty")
	}

	res := []string{}
	j := 0

	for i := 0; i < len(s)-len(sep); {
		if s[i:i+len(sep)] == sep {
			res = append(res, s[j:i])
			j = i
		} else {
			i++
		}
	}
	res = append(res, s[j:])
	return res, nil
}
