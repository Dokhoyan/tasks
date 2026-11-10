package main

import "sync"

//Нужно выполнить все checks параллельно и вернуть, набралось ли хотя бы required ответов true.
//Даже при уже известном результате нужно дождаться всех; required = 0 не отменяет вызовы.
//Проверок не более 32, 0 <= required <= len(checks).

func CheckQuorum(checks []func() bool, required int) bool {
	trueCnt := 0

	mu := sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(len(checks))
	for _, check := range checks {
		go func() {
			defer wg.Done()
			if check() {
				mu.Lock()
				trueCnt++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	return trueCnt >= required
}
