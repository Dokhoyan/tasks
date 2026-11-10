package main

//Нужно сохранить первые k чисел из канала in, затем считать и отбросить остальные до закрытия
//канала отправителем. Если чисел меньше k, нужно вернуть все. k >= 0; при k = 0 тоже нужно
//дочитать канал.

func FirstK(in <-chan int, k int) []int {
	res := make([]int, 0, k)

	for n := range in {
		if k != 0 {
			res = append(res, n)
			k--
		}
	}
	return res
}