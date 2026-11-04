package main

func Contains(s, sub string) bool {
	if len(sub) == 0 {
		return true
	}
	if len(sub) > len(s) {
		return false
	}

	for i := 0; i <= len(s)-len(sub); i++ {
		j := 0
		for j < len(sub) && sub[j] == s[i+j] {
			j++
		}

		if j == len(sub) {
			return true
		}
	}
	return false
}
