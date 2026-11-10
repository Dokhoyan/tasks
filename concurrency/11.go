package main

import "sync"

//Нужно запустить переданные функции-проверки first и second в двух горутинах и дождаться
//обеих. Ошибка одной не отменяет вторую; нужно вернуть ошибки в порядке first, second.

func RunChecks(first, second func() error) (error, error) {
	wg := sync.WaitGroup{}
	var f error
	var s error

	wg.Add(2)
	go func() {
		defer wg.Done()
		f = first()
	}()

	go func() {
		defer wg.Done()
		s = second()
	}()
	wg.Wait()
	return f, s
}
