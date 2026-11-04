package main

func ConcatMany(parts ...[]int) []int {
	resLen := 0
	for _, part := range parts {
		resLen += len(part)
	}
	res := make([]int, resLen)
	i := 0
	for _, part := range parts {
		copy(res[i:], part)
		i += len(part)
	}

	return res
}
