package main

func Unique(s []int) []int {
	res := make([]int, 0)
	unique := make(map[int]struct{})

	for _, v := range s {
		if _, ok := unique[v]; !ok {
			res = append(res, v)
			unique[v] = struct{}{}
		}
	}
	
	return res
}
