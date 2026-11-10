package main

import "sync"

//Нужно выполнить first и second параллельно, затем вызвать finish один раз. Сразу нужно вернуть
//канал-сигнал и закрыть его после завершения всех трех вызовов.

func RunPairThenFinish(first, second, finish func()) <-chan struct{}{
	signal := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {

		go func() {
			defer wg.Done()
			first()
		}()

		go func() {
			defer wg.Done()
			second()
		}()
		wg.Wait()
		finish()
		close(signal)
	}()

	return signal
}