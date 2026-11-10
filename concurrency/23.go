package main

// Нужно выполнить calculate в фоне и сразу вернуть канал одного ответа. При ошибке нужно
// отправить Value = 0 и исходный Err, при успехе сохранить значение. Отправка и закрытие выхода
// не должны ждать читателя.
type Outcome struct {
	Value int
	Err   error
}

func StartCalculation(calculate func() (int, error)) <-chan Outcome {
	ch := make(chan Outcome, 1)

	go func() {
		defer close(ch)
		v, err := calculate()
		if err != nil {
			ch <- Outcome{0, err}
		} else {
			ch <- Outcome{v, err}
		}
	}()
	
	return ch
}
