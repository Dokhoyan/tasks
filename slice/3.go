package main

import "math"

func Max(s []int) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	max := math.MinInt64
	for _, v := range s {
		if v > max {
			max = v
		}
	}

	return max, true
}
