package main

func Clone(m map[string][]int) map[string][]int {
	res := make(map[string][]int, len(m))

	for k, v := range m {
		if v == nil {
			res[k] = nil
			continue
		}

		v2 := make([]int, len(v))
		copy(v2, v)
		res[k] = v2
	}
	return res
}
