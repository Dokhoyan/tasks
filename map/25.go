package main

func FilterMap(m map[string]int, keep func(string, int) bool) map[string]int {
	for k, v := range m {
		if !keep(k, v) {
			delete(m, k)
		}
	}
	return m
}
