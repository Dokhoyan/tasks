package main

import (
	"errors"
	"fmt"
)

func RemoveAt(s []int, i int) ([]int, error) {
	if i < 0 || i >= len(s) {
		return nil, errors.New("index out of range")
	}

	res := append(s[:i], s[i+1:]...)
	fmt.Println(s)
	return res, nil
}

func main() {
	s := []int{1, 2, 3}
	a, _ := RemoveAt(s, 2)
	fmt.Println(a)
	a[0] = 0
	fmt.Println(a)
	fmt.Println(s)
}
