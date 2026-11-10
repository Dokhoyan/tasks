package main

import (
	"errors"
	"fmt"
)

// Нужно считать n частей из канала in и объединить их Values в порядке Index от 0 до n-1. Пропуск,
// повтор или неверный индекс дают nil и ошибку; в любом случае нужно дочитать канал до
// закрытия отправителем. n >= 0, пустая часть учитывается; отправитель не меняет Values после
// передачи, результат не разделяет с ними элементы.
type ResultPart struct {
	Index  int
	Values []int
}

func JoinParts(in <-chan ResultPart, n int) ([]int, error) {
	parts := make([][]int, n)
	seen := make([]bool, n)

	var firstErr error
	for r := range in {
		if r.Index < 0 || r.Index >= n {
			if firstErr == nil {
				firstErr = errors.New("index out of range")
			}
			continue
		}

		if seen[r.Index] {
			if firstErr == nil {
				firstErr = errors.New("duplicate index")
			}
			continue
		}

		parts[r.Index] = append(parts[r.Index], r.Values...)
		seen[r.Index] = true
	}

	if firstErr != nil {
		return nil, firstErr
	}

	for b := range seen {
		if !seen[b] {
			return nil, fmt.Errorf("index out of range")
		}
	}

	result := make([]int, 0)
	for _, part := range parts {
		for _, value := range part {
			result = append(result, value)
		}
	}
	return result, nil
}

//Index: 2, Values: [40, 50]
//Index: 0, Values: [10]
//Index: 1, Values: [20, 30]
//
//Результат: [10, 20, 30, 40, 50]
