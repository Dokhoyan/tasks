package main

func ContainsAll(have, need []string) bool {
	haveMap := make(map[string]struct{}, len(have))
	for _, v := range have {
		haveMap[v] = struct{}{}
	}

	for _, v := range need {
		if _, ok := haveMap[v]; !ok {
			return false
		}
	}

	return true
}
