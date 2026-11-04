package main

func FirstUniqueRune(s string) (rune, bool) {
	found := make(map[rune]int, len(s))

	for _, r := range s {
		found[r]++
	}

	for _, r := range s {
		if found[r] == 1 {
			return r, true
		}
	}

	return 0, false
}
