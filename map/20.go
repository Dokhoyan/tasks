package main

import (
	"sort"
	"unicode"
)

func TopKWords(text string, k int) []string {
	if k == 0 {
		return nil
	}

	setWords := make(map[string]int, k)
	word := []rune{}
	words := make([]string, 0)

	for _, v := range text {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			word = append(word, v)
		} else if len(word) > 0 {
			setWords[string(word)]++
			word = word[:0]
		}
	}

	if len(word) > 0 {
		setWords[string(word)]++
	}

	for k := range setWords {
		words = append(words, k)
	}

	sort.Slice(words, func(i, j int) bool {
		if setWords[words[i]] != setWords[words[j]] {
			return setWords[words[i]] > setWords[words[j]]
		}
		return words[i] < words[j]
	})

	if k > len(words) {
		k = len(words)
	}

	return words[:k]
}
