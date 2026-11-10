package main

import "sync"

//Нужно запустить jobs параллельно и сразу вернуть небуферизированный канал ответов с
//индексами работ в порядке получения. Нужно передать ровно один ответ на работу и закрыть
//канал после всех работ. Работ не более 32; получатель прочитает все ответы до закрытия, нельзя
//накапливать их все перед первой отправкой.

type IndexedValue struct { Index, Value int }

func StartIndexed(jobs []func() int) <-chan IndexedValue{
	ch := make(chan IndexedValue)

	wg := sync.WaitGroup{}
	wg.Add(len(jobs))
	for i, job := range jobs {
		go func() {
			defer wg.Done()
			val := job()
			ch <- IndexedValue{Index: i, Value: val}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}