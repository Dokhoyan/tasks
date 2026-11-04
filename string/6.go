package main

func RemoveRune(s string, target rune) string {
	runes := []rune(s)
	res := make([]rune, 0, len(runes))

	for _, v := range runes {
		if target != v {
			res = append(res, v)
		}

	}

	return string(res)
}
