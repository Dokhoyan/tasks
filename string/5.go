package main

import "fmt"

func Reverse(s string) string {
	runc := []rune(s)
	res := make([]rune, 0, len(runc))

	for i := len(runc) - 1; i >= 0; i-- {
		res = append(res, runc[i])
	}
	return string(res)
}

func main() {
	fmt.Println(Reverse("12345"))
}
