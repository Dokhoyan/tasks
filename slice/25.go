package main

func Clone(s []int) []int {
	if s == nil {
		return nil
	}
	
	res := make([]int, len(s))
	copy(res, s)
	return res
}
