package main

//Нужно запустить переданную функцию compute в отдельной горутине и сразу вернуть канал для
//результата. Когда вычисление закончится, нужно отправить полученное число в канал и закрыть
//его, не ожидая читателя.

func StartValue(compute func() int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		res := compute()
		ch <- res
		close(ch)
	}()

	return ch
}
