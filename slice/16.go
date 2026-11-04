package main

func RemoveAllInPlace(s []int, target int) []int {
	i := 0
	for _, v := range s {
		if v != target {
			s[i] = v
			i++
		}
	}
	return s[:i]
}
