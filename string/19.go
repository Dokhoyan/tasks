package main

func TrimSet(s, cutset string) string {
	if len(s) == 0 {
		return ""
	}

	setCutset := make(map[rune]struct{}, len(cutset))
	for _, v := range []rune(cutset) {
		setCutset[v] = struct{}{}
	}
	runeS := []rune(s)
	left, right := 0, len(runeS)
	for left < right {
		if _, ok := setCutset[runeS[left]]; !ok {
			break
		}
		left++
	}
	for left < right {
		if _, ok := setCutset[runeS[right-1]]; !ok {
			break
		}
		right--
	}
	return string(runeS[left:right])
}
