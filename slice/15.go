package main

func RemoveAll(s []int, target int) []int {
	new := []int{}
	for _, v := range s {
		if v != target {
			new = append(new, v)
		}
	}
	return new
}
