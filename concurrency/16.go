package main

//Нужно считывать строки из канала in до его закрытия отправителем и посчитать, сколько раз
//встретилась каждая. Регистр и пустая строка значимы; если строк не было, нужно вернуть пустую
//мапу.

func WordCounts(in <-chan string) map[string]int {
	uniqueStr := make(map[string]int)

	for str := range in {
		uniqueStr[str] += 1
	}
	return uniqueStr
}
