package main

func UniqueInPlace(s []int) []int {
	unique := make(map[int]struct{})

	uniqueLen := 0
	l := len(s)
	for i := 0; i < l; i++ {
		if _, ok := unique[s[i]]; !ok {
			unique[s[i]] = struct{}{}
			s[uniqueLen] = s[i]
			uniqueLen++
		}
	}
	return s[:uniqueLen]
}
