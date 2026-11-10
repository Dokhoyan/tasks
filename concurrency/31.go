package main

import "sync"

//Параллельно нужно вызвать analyze для каждой части и вернуть ответы в порядке parts,
//дождавшись всех вызовов. Нельзя пропускать пустые и повторяющиеся части. Частей не более
//32; данные никто не меняет, analyze только читает их.

func AnalyzeParts(parts [][]int, analyze func([]int) int) []int {
	res := make([]int, len(parts))

	wg := new(sync.WaitGroup)
	wg.Add(len(parts))
	for i, part := range parts {
		go func(part []int, i int) {
			defer wg.Done()
			res[i] = analyze(part)
		}(part, i)
	}

	wg.Wait()
	return res
}
