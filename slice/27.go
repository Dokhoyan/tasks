package main

func Concat(a, b []int) []int {
	res := make([]int, len(a)+len(b))
	copy(res, a)
	copy(res[len(a):], b)
	return res
}
