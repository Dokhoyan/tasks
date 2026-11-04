package main

import "unicode"

func AreAnagrams(a, b string) bool {
	setA := make(map[rune]int, len(a))

	for _, v := range a {
		if unicode.IsSpace(v) {
			continue
		}
		v := unicode.ToLower(v)
		setA[v]++
	}

	for _, v := range b {
		if unicode.IsSpace(v) {
			continue
		}
		v := unicode.ToLower(v)

		if setA[v] == 0 {
			return false
		}

		setA[v]--
		if setA[v] == 0 {
			delete(setA, v)
		}
	}

	return len(setA) == 0
}
