package main

func ReplaceRange(s []int, left, right int, values []int) ([]int, error) {
	oldLen := len(s)
	newLen := oldLen - (right - left) + len(values)

	if newLen > oldLen {
		s = append(s, make([]int, newLen-oldLen)...)
	}

	copy(s[left+len(values):], s[right:oldLen])

	copy(s[left:], values)

	return s[:newLen], nil
}
