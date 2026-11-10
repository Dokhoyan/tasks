package main

import "sync"

//Параллельно нужно найти первое вхождение target в двух соседних половинах nums. Нужно
//дождаться обеих работ и вернуть минимальный индекс во всем слайсе или -1, если числа нет.
//nums нельзя менять.

func FindFirst(nums []int, target int) int {
	mid := len(nums) / 2
	left, right := -1, -1

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < mid; i++ {
			if nums[i] == target {
				left = i
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := mid; i < len(nums); i++ {
			if nums[i] == target {
				right = i
				return
			}
		}
	}()

	wg.Wait()

	if left != -1 {
		return left
	}
	return right
}
