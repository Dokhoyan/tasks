package main

//Нужно считывать числа из in до закрытия канала отправителем. Для каждого последовательно
//нужно вызвать check, а при успехе - apply; нужно вернуть ошибку check либо результат apply на
//соответствующей позиции. Нельзя останавливать обработку следующих чисел после ошибок.

func CheckThenApply(in <-chan int, check, apply func(int) error) []error{
	errors := make([]error, 0)

	for n := range in {
		err := check(n)
		if err == nil {
			err = apply(n)
		}
		errors = append(errors, err)
	}
	return errors
}