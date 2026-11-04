package main

import "fmt"

func IsASCII(s string) bool {
	return len(s) == len([]rune(s))
}

func main() {
	fmt.Println(IsASCII("as"))
	fmt.Println(IsASCII("ыв"))
}
