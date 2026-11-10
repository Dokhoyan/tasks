package main

import "sync"

//Параллельно нужно вызвать left и right с независимыми копиями nums и вернуть измененные
//копии в этом порядке после обоих вызовов. Исходный слайс нельзя менять; результаты не
//должны разделять элементы друг с другом или с ним. Обе функции вызываются и для пустого
//входа; они не сохраняют слайсы.

func ForkEdits(nums []int, left, right func([]int)) ([]int, []int) {
	numsL := make([]int, len(nums))
	numsR := make([]int, len(nums))
	copy(numsL, nums)
	copy(numsR, nums)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		left(numsL)
	}()
	go func() {
		defer wg.Done()
		right(numsR)
	}()

	wg.Wait()
	return numsL, numsR
}
