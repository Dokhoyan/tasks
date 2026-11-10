package main

import "fmt"

//Нужно считывать числа из каналов left и right и объединять их по позициям: первое с первым,
//второе со вторым. Нужно дочитать оба канала до закрытия отправителями; при разной длине
//нужно вернуть полные пары и ошибку, иначе - пары и nil. Каналы разные, отправители
//независимы; два канала без значений допустимы.

type NumberPair struct{ Left, Right int }

func PairStreams(left, right <-chan int) ([]NumberPair, error) {
	res := make([]NumberPair, 0)

	for {
		l, okL := <-left
		r, okR := <-right

		if okL && okR {
			res = append(res, NumberPair{Left: l, Right: r})
			continue
		}

		if !okL && !okR {
			return res, nil
		}

		if okL {
			for range left {
			}
		} else {
			for range right {
			}
		}

		return res, fmt.Errorf("streams have different lengths")
	}
}
