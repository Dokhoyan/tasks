package main

//Параллельно нужно посчитать частоты чисел в частях chunks и объединить результаты в одну
//мапу. Итоговую мапу изменяет один получатель; переданные ему локальные мапы отправители
//больше не используют. Частей не более 32, пустые тоже дают результат; chunks нельзя менять.

func CountChunks(chunks [][]int) map[int]int {
	res := make(map[int]int)
	ch := make(chan map[int]int, len(chunks))
	defer close(ch)

	for _, chunk := range chunks {
		go func() {
			set := make(map[int]int)
			for _, v := range chunk {
				set[v]++
			}
			ch <- set
		}()
	}

	for range chunks {
		local := <-ch
		for key, count := range local {
			res[key] += count
		}
	}

	return res
}
