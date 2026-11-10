package main

import "sync"

//Нужно запустить background в новой горутине, а local выполнить в вызывающей горутине
//параллельно с ней. Нужно дождаться обеих и вернуть результаты в порядке local, background.

func RunAlongside(local, background func() int) (int, int) {
	l, b := 0, 0

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		b = background()
	}()

	l = local()
	wg.Wait()

	return l, b
}
