package main

//В одной фоновой горутине нужно последовательно вычислить next(0), ..., next(n-1) и отправить
//ответы в небуферизированный канал. Нужно вернуть канал сразу и закрыть после отправок; при
//n = 0 вызовов нет. n >= 0, получатель дочитает выход.

func Generate(n int, next func(int) int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- next(i)
		}
	}()

	return ch
}
