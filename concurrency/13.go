package main

//Нужно запустить горутину, которая читает числа из in и для каждого по очереди вызывает
//переданную функцию consume. Сразу нужно вернуть канал-сигнал и закрыть его после закрытия
//in и завершения всех вызовов consume. Если чисел не было, сигнал тоже должен закрыться.

func StartConsumer(in <-chan int, consume func(int)) <-chan struct{} {
	signal := make(chan struct{})

	go func() {
		for v := range in {
			consume(v)
		}
		close(signal)
	}()

	return signal
}
