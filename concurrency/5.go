package main

func PrepareNumbers(nums []int) <-chan int {
	ch := make(chan int, len(nums))
	defer close(ch)

	for _, v := range nums {
		ch <- v
	}
	return ch
}
