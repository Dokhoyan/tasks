package main

func Sum(s []int) int {
	sum := 0

	for _, v := range s {
		sum += v
	}

	return sum
}
