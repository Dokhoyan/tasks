package main

import "sort"

//"кот" и "ток" — анаграммы.
//"сон" и "нос" — анаграммы.
//"кот" и "кит" — нет: буквы отличаются.
//"ааб" и "абб" — нет: количество каждой буквы отличается.

func GroupAnagrams(words []string) [][]string {
	res := make([][]string, 0)
	setAnagrams := make(map[string]int)

	i := 0
	for _, word := range words {
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] > runes[j]
		})
		if _, ok := setAnagrams[string(runes)]; !ok {
			setAnagrams[string(runes)] = i
			i++
			res = append(res, []string{})
		}

	}

	for _, word := range words {
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] > runes[j]
		})
		res[setAnagrams[string(runes)]] = append(res[setAnagrams[string(runes)]], word)
	}

	return res
}
