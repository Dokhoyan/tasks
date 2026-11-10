package main

//Нужно считать числа из in до закрытия канала отправителем и вернуть последние k в исходном
//порядке. Нужно использовать O(k) памяти вместе с результатом. k >= 0; при k = 0 тоже нужно
//дочитать канал и вернуть пустой слайс.

func LastK(in <-chan int, k int) []int {
	res := make([]int, 0, k)

	for n := range in {
		if k == 0 {
			continue
		}

		if len(res) == k {
			copy(res[:len(res)-1], res[1:])
			res[k-1] = n
		} else {
			res = append(res, n)
		}
	}

	return res
}
