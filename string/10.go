package main

import "unicode"

func NormalizeSpaces(s string) string {
	runs := []rune(s)
	res := make([]rune, 0, len(runs))
	var isSpace bool

	for i := range runs {
		if unicode.IsSpace(runs[i]) {
			isSpace = len(res) > 0
			continue
		}

		if isSpace {
			res = append(res, ' ')
			isSpace = false
		}
		
		res = append(res, runs[i])
	}

	return string(res)
}
