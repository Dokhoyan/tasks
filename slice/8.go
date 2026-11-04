package main

func IndexesOf(s []int, target int) []int {
	res := []int{}

	for i, v := range s {
		if v == target {
			res = append(res, i)
		}
	}
	return res
}
