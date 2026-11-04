package main

import "fmt"

func ReplaceRune(s string, old, new rune) string {
	runes := []rune(s)
	for i, v := range runes {
		if v == old {
			runes[i] = new
		}
	}
	return string(runes)
}

func main() {
	fmt.Println(ReplaceRune("123345", '3', '7'))
}
