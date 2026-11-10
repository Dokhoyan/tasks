package main

//Нужно считывать числа из канала in до его закрытия отправителем и проверить, что каждое
//следующее не меньше предыдущего. После нарушения тоже нужно дочитать канал. Если чисел
//нет или оно одно, нужно вернуть true.

func IsNonDecreasing(in <-chan int) bool {
	flag := true
	prev := <-in

	for v := range in {
		if v < prev {
			flag = false
		}
		prev = v
	}

	return flag
}
