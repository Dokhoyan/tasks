package main

func GroupByParity(nums []int) map[string][]int {
	res := map[string][]int{
		"even": nil,
		"odd":  nil,
	}

	for _, num := range nums {
		if num%2 == 0 {
			res["even"] = append(res["even"], num)
		} else {
			res["odd"] = append(res["odd"], num)
		}
	}
	return res
}
