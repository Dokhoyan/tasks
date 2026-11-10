//package main
//
//// В одной горутине нужно считывать числа из in и последовательно вызывать transform. Сразу
//// нужно вернуть канал ответов: один на число, при ошибке Value = 0 и исходный Err. После
//// закрытия in отправителем нужно закрыть ответы; нельзя накапливать их все перед отправкой,
//// получатель читает до закрытия.
//type Outcome struct {
//	Value int
//	Err   error
//}
//
//func TransformStream(in <-chan int, transform func(int) (int, error)) <-chan Outcome {
//
//}
