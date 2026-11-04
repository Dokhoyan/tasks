package main

func IndexOf(s, sub string) int {
	if len(sub) == 0 {
		return -1
	}
	if len(sub) > len(s) {
		return -1
	}

	for i := 0; i <= len(s)-len(sub); i++ {
		j := 0
		for j < len(sub) && sub[j] == s[i+j] {
			j++
		}

		if j == len(sub) {
			return i
		}
	}
	return -1
}
