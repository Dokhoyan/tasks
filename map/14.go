package main

import "sort"

func SymmetricDifference(a, b []int) []int {
	res := []int{}
	aMap := make(map[int]struct{}, len(b)+len(a))
	bMap := make(map[int]struct{}, len(b)+len(a))

	for _, v := range a {
		aMap[v] = struct{}{}
	}

	for _, v := range b {
		bMap[v] = struct{}{}
	}

	for k := range aMap {
		if _, ok := bMap[k]; !ok {
			res = append(res, k)
		}
	}

	for k := range bMap {
		if _, ok := aMap[k]; !ok {
			res = append(res, k)
		}
	}

	sort.Ints(res)
	return res
}
