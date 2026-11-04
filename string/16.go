package main

import "fmt"

func LongestCommonPrefix(items []string) string {
	if len(items) == 0 {
		return ""
	}

	prefix := []rune(items[0])

	for _, item := range items[1:] {
		runeItem := []rune(item)
		j := 0
		for j < len(runeItem) && j < len(prefix) {
			if runeItem[j] != prefix[j] {
				break
			}
			j++
		}
		prefix = prefix[:j]
	}
	return string(prefix)
}

func main() {
	fmt.Println(LongestCommonPrefix([]string{"asdfff", "asdrfr", "asdeer"}))
}
