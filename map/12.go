package main

func Difference(a, b []int) []int {
	res := []int{}
	bMap := make(map[int]struct{}, len(b))
	aMap := make(map[int]struct{}, len(a))
	for _, v := range b {
		bMap[v] = struct{}{}
	}

	for _, v := range a {
		if _, ok := bMap[v]; !ok {
			if _, ok := aMap[v]; !ok {
				res = append(res, v)
				aMap[v] = struct{}{}
			}
		}
	}
	return res
}
