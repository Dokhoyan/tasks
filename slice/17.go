package main

func KeepEvenInPlace(s []int) []int {
	i := 0

	for v := range s {
		if v%2 == 0 {
			s[i] = v
			i++
		}
	}
	
	return s[:i+1]
}
