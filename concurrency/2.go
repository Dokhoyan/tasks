package main

func SumStream(in <-chan int) int {
	sum := 0
	for n := range in {
		sum += n
	}
	return sum
}
