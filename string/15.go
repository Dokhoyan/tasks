package main

func CommonPrefix(a, b string) string {
	reueA := []rune(a)
	reueB := []rune(b)

	i := 0
	for i = 0; i < len(reueA) && i < len(reueB); i++ {
		if reueA[i] == reueB[i] {
			continue
		}
		break
	}
	return string(reueA[:i])
}
