package main

import "fmt"

func StringSize(s string) (bytes int, runes int) {
	return len(s), len([]rune(s))
}

func main() {
	fmt.Println(StringSize("asdфfg"))
}
