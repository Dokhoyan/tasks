package main

func Compose(first map[string]string, second map[string]int) map[string]int {
	res := make(map[string]int)
	for k, _ := range first {
		if _, ok := second[first[k]]; ok {
			res[k] = second[first[k]]
		}
	}
	return res
}
