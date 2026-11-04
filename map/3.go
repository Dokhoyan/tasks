package main

func Unique(nums []int) []int {
	res := []int{}
	uniqueCheck := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := uniqueCheck[v]; !ok {
			res = append(res, v)
			uniqueCheck[v] = struct{}{}
		}
	}

	return res
}
