package main

func MapValues(m map[string]int, f func(int) int) map[string]int {
	res := make(map[string]int, len(m))

	for key, value := range m {
		res[key] = f(value)
	}

	return res
}
