package main

//Нужно выполнить work в фоне и сразу вернуть канал-сигнал без значений. Его закрытие после
//work должны увидеть все наблюдатели, включая начавших ждать после завершения.

func StartCompletion(work func()) <-chan struct{} {
	signalCh := make(chan struct{})

	go func() {
		defer close(signalCh)
		work()
	}()

	return signalCh
}