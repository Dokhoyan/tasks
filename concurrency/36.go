package main

import "sync"

//Нужно запустить проверки независимо: каждая ждет закрытия своего Ready, затем вызывает
//Check. Закрытый Ready поздней проверки не должен ждать разрешения ранней. Нужно
//дождаться всех и вернуть ошибки по индексам, включая nil; проверок не более 32, их владельцы
//гарантированно закрывают Ready без отправки значений.
type GatedCheck struct { Ready <-chan struct{}; Check func() error }

func RunReadyChecks(checks []GatedCheck) []error {
	resErrs := make([]error, len(checks))

	wg := new(sync.WaitGroup)
	wg.Add(len(checks))
	for i, check := range checks {
		go func() {
			defer wg.Done()
			<-check.Ready
			resErrs[i] = check.Check()
		}()
	}

	wg.Wait()
	return resErrs
}