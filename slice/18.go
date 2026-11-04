package main

func MoveZerosToEnd(s []int) {
	i := 0

	for _, v := range s {
		if v != 0 {
			s[i] = v
			i++
		}
	}

	for ind := i; ind < len(s); ind++ {
		s[ind] = 0
	}

}
