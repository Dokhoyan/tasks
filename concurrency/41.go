package main

//Нужно считывать строки из in до закрытия канала отправителем и последовательно вызывать
//apply для каждой, продолжая после ошибок. Нужно вернуть по одному error на строку, включая nil
//для успешных вызовов.

func ApplyAll(in <-chan string, apply func(string) error) []error {
	res := make([]error, 0)

	for v := range in {
		res = append(res, apply(v))
	}
	return res
}
