package main

func MergeCounts(a, b map[string]int) map[string]int {
	res := make(map[string]int)

	for k, _ := range a {
		if _, ok := b[k]; ok {
			res[k] = a[k] + b[k]
		}
	}

	return res
}
