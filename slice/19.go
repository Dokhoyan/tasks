package main

func CompactNonZero(s []int) []int {
	i := 0

	for _, v := range s {
		if v != 0 {
			s[i] = v
			i++
		}
	}

	return s[:i]

}
