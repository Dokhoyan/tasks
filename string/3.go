package main

func FirstRune(s string) (rune, bool) {
	if len(s) == 0 {
		return 0, false
	}
	for _, v := range s {
		return v, true
	}
	return 0, false
}
