package main

import (
	"strconv"
	"strings"
)

func EncodeRLE(s string) string {
	if s == "" {
		return ""
	}

	runes := []rune(s)
	res := strings.Builder{}
	prev := runes[0]
	cnt := 1
	for _, v := range runes[1:] {
		if v == prev {
			cnt++
		} else {
			res.WriteString(strconv.Itoa(cnt))
			res.WriteString(":")
			res.WriteString(string(prev))
			res.WriteString(";")
			prev = v
			cnt = 1
		}

	}
	res.WriteString(strconv.Itoa(cnt))
	res.WriteString(":")
	res.WriteString(string(prev))

	return res.String()
}
