package main

import "fmt"

func ReadExactlyTwo(in <-chan int) (int, int, error) {
	first := 0
	second := 0
	cnt := 0
	for n := range in {
		if cnt == 0 {
			first = n
		} else {
			second = n
		}
		cnt++
	}
	if cnt != 2 {
		return 0, 0, fmt.Errorf("ReadExactlyTwo: expected 2 got %d", cnt)
	}

	return first, second, nil
}
