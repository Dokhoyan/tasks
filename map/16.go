package main

func MostFrequent(nums []int) (value int, count int, ok bool) {
	if len(nums) == 0 {
		return 0, 0, false
	}

	numsCnt := make(map[int]int, len(nums))

	for _, num := range nums {
		numsCnt[num]++
		if numsCnt[num] > count || (numsCnt[num] == count && num < value) {
			count = numsCnt[num]
			value = num
		}
	}

	return value, count, true
}
