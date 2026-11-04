package main

func LastRune(s string) (rune, bool) {
	if len(s) == 0 {
		return 0, false
	}
	runs := []rune(s)
	return runs[len(runs)-1], true
}
