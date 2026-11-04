package main

import "errors"

func CountOccurrences(s, sub string) (int, error) {
	if len(sub) == 0 {
		return -1, errors.New("sub is empty")
	}

	subCnt := 0
	for i := 0; i <= len(s)-len(sub); {
		j := 0
		for j < len(sub) && s[i+j] == sub[j] {
			j++
		}

		if j == len(sub) {
			subCnt++
			i += len(sub)
			continue
		}
		i++
	}
	return subCnt, nil
}
