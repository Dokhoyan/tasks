package main

//Нужно запустить горутину, которая читает числа из in и отправляет отрицательные в один
//небуферизированный канал, остальные - в другой. Сразу нужно вернуть эти два канала в
//указанном порядке; нужно сохранить порядок чисел в каждом и закрыть оба после закрытия in
//отправителем. Получатели одновременно читают оба канала до закрытия.

func RouteBySign(in <-chan int) (<-chan int, <-chan int) {
	otr := make(chan int)
	notr := make(chan int)

	go func() {
		defer close(notr)
		defer close(otr)
		for n := range in {
			if n < 0 {
				otr <- n
			} else {
				notr <- n
			}
		}
	}()

	return otr, notr
}
