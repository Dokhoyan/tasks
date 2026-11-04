package main

import "math"

func Min(s []int) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	min := math.MaxInt64
	for _, v := range s {
		if v < min {
			min = v
		}
	}

	return min, true
}
