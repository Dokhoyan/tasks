package main

import (
	"fmt"
	"strings"
)

func CollapseRuns(s string) string {
	runeS := []rune(s)
	res := strings.Builder{}
	first := true
	var prev rune

	for _, v := range runeS {
		if first || v != prev {
			res.WriteRune(v)
		}
		prev = v
		first = false
	}
	return res.String()
}

func main() {
	fmt.Println(CollapseRuns("ааа"))
	fmt.Println(CollapseRuns("абв"))
	fmt.Println(CollapseRuns("абвг"))
}
