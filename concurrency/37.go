package main

import (
	"errors"
	"sync"
)

//Нужно параллельно вызвать left и right и дождаться обеих операций. При равных строках без
//ошибок нужно вернуть строку; при различии - ошибку. Ошибки источников имеют приоритет над
//сравнением, left - над right; при любой ошибке строковый результат пустой.

func CompareSources(left, right func() (string, error)) (string, error) {
	var leftValue, rightValue string
	var leftErr, rightErr error

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		leftValue, leftErr = left()
	}()

	go func() {
		defer wg.Done()
		rightValue, rightErr = right()
	}()

	wg.Wait()

	if leftErr != nil {
		return "", leftErr
	}
	if rightErr != nil {
		return "", rightErr
	}
	if leftValue != rightValue {
		return "", errors.New("sources returned different strings")
	}

	return leftValue, nil
}
