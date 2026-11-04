package main

import (
	"strings"
)

func UniqueFold(words []string) []string {
	res := []string{}
	wordsSet := map[string]struct{}{}

	for _, w := range words {
		wLowwer := strings.ToLower(w)
		if _, ok := wordsSet[wLowwer]; !ok {
			res = append(res, w)
			wordsSet[wLowwer] = struct{}{}
		}
	}

	return res
}
