package main

func CountEven(s []int) int {
	chetCnt := 0
	for _, v := range s {
		if v%2 == 0 {
			chetCnt++
		}
	}
	return chetCnt
}
