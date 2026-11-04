package main

func Intersection(a, b []string) []string {
	res := []string{}
	bMap := make(map[string]struct{}, len(b))
	aMap := make(map[string]struct{}, len(a))

	for _, v := range b {
		bMap[v] = struct{}{}
	}

	for _, v := range a {
		if _, ok := bMap[v]; ok {
			if _, ok := aMap[v]; ok {
				continue
			}
			res = append(res, v)
			aMap[v] = struct{}{}
		}
	}

	return res
}