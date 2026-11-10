package main

func MinStream(in <-chan int) (int, bool) {
	var min int
	found := false
	for v := range in {
		if !found || v < min {
			min = v
		}
		found = true
	}
	return min, found
}
