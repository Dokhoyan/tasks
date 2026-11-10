package main

//Нужно запустить горутину, которая отправляет числа nums в канал в исходном порядке, затем
//закрывает его. Сразу нужно вернуть этот канал; его емкость равна capacity. capacity >= 0,
//получатель прочитает все значения до закрытия; nums никто не меняет до завершения отправки.

func BufferedSource(nums []int, capacity int) <-chan int {
	ch := make(chan int, capacity)

	go func() {
		defer close(ch)
		for _, n := range nums {
			ch <- n
		}
	}()

	return ch
}
