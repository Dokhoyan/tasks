package main

//Нужно считывать числа из in и последовательно вызывать apply до первой ошибки. После нее
//больше нельзя вызывать apply, но нужно дочитать канал до закрытия отправителем. Нужно
//вернуть первую ошибку или nil.

func ApplyUntilError(in <-chan int, apply func(int) error) error {
	var err error
	for n := range in {
		if err == nil {
			err = apply(n)
		}
	}
	return err
}
