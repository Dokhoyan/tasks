package main

import (
	"fmt"
	"unicode"
)

func WordFrequency(text string) map[string]int {
	setWords := make(map[string]int)
	word := []rune{}
	for _, v := range text {

		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			word = append(word, unicode.ToLower(v))
		} else if len(word) > 0 {
			setWords[string(word)]++
			word = word[:0]
		}
	}

	if len(word) > 0 {
		setWords[string(word)]++
	}
	return setWords
}

func main() {
	fmt.Println(WordFrequency("Go go GO"))
}
