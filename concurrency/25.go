package main

import (
	"fmt"
)

//Нужно считать из канала in длину пакета и указанное число элементов; затем канал должен
//закрыться. Отсутствие заголовка, отрицательная длина, недостача или избыток элементов дают
//nil и ошибку; в любом случае нужно дочитать канал до закрытия отправителем. Неотрицательная
//длина не больше 1024; длина 0 задает пустой пакет.

func ReadFrame(in <-chan int) ([]int, error) {
	length, ok := <-in
	if !ok {
		return nil, fmt.Errorf("missing header")
	}

	if length < 0 {
		for range in {
		}
		return nil, fmt.Errorf("negative length: %d", length)
	}

	res := make([]int, 0, length)
	count := 0

	for v := range in {
		if count < length {
			res = append(res, v)
		}
		count++
	}

	if count != length {
		return nil, fmt.Errorf(
			"expected %d elements, got %d", length, count,
		)
	}

	return res, nil
}
