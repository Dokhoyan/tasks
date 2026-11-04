package main

func LastIndex(words []string) map[string]int {
	res := make(map[string]int)

	for i, v := range words {
		res[v] = i
	}

	return res
}
