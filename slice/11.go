package main

func SwapPairs(s []int) {
	left, right := 0, 1

	for right < len(s) {
		s[left], s[right] = s[right], s[left]
		left += 2
		right += 2
	}
}
