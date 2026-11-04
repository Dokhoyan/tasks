package main

func CountOccurrences(s []int, target int) int {
	targetCnt := 0
	for _, v := range s {
		if v == target {
			targetCnt++
		}
	}
	return targetCnt
}
