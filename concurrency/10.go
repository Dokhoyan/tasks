package main

import "sync"

//Нужно параллельно посчитать вхождения target в left и right, дождаться обеих работ и вернуть
//количества в порядке аргументов. Для пустого слайса нужно вернуть 0. Входные данные менять
//нельзя.

func CountInBoth(left, right []int, target int) (int, int) {
	l, r := 0, 0
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, v := range left {
			if v == target {
				l++
			}
		}
	}()

	go func() {
		defer wg.Done()
		for _, v := range right {
			if v == target {
				r++
			}
		}
	}()
	wg.Wait()
	return l, r
}
