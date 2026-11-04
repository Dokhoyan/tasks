package main

func CountInts(nums []int) map[int]int {
	res := make(map[int]int)
	for _, v := range nums {
		res[v]++
	}
	return res
}
