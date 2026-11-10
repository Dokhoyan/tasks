package main

import "fmt"

func KeepEvenInPlace(s []int) []int {
	//if s == nil {
	//	return nil
	//}
	i := 0

	for _, v := range s {
		fmt.Println(v)
		if v%2 == 0 {
			s[i] = v
			i++
		}
	}

	return s[:i]
}

func main() {
	fmt.Println(KeepEvenInPlace(nil))
}
