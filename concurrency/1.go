package main

func ReceiveN(in <-chan int, n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = <-in
	}

	return res
}
