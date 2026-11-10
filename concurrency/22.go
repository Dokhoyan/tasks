package main

import "sync"

//Нужно запустить все переданные функции jobs параллельно и сразу вернуть общий канал их
//результатов в порядке получения. Емкость канала равна len(jobs): работы и закрытие канала
//должны завершиться без читателя. Не более 32 работ; пустой список дает закрытый канал.

func StartBatch(jobs []func() int) <-chan int {
	ch := make(chan int, len(jobs))

	wg := new(sync.WaitGroup)
	for _, job := range jobs {
		wg.Add(1)
		go func(job func() int) {
			defer wg.Done()
			ch <- job()
		}(job)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}
