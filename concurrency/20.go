package main

//Сразу нужно вернуть канал, а read вызвать в фоне только после закрытия ready. Нужно отправить
//один ответ и закрыть выход без ожидания читателя. Владелец гарантированно закроет ready без
//отправки значений.

func ReadAfterReady(ready <-chan struct{}, read func() int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		defer close(ch)
		<-ready
		ch <- read()
	}()

	return ch
}
