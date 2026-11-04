package main

func ReplaceAll(s []int, old, new int) {
	for i, v := range s {
		if v == old {
			s[i] = new
		}
	}
}
