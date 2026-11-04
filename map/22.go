package main

func GroupByLength(words []string) map[int][]string {
	res := make(map[int][]string)

	for _, word := range words {
		runes := []rune(word)
		res[len(runes)] = append(res[len(runes)], word)
	}

	return res
}
