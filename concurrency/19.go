package main

//Нужно дождаться закрытия всех каналов signals. Владельцы гарантированно закрывают их без
//отправки значений; каналы не nil, могут повторяться и быть уже закрыты. Пустой список не
//требует ожидания.

func AwaitSignals(signals []<-chan struct{}) {
	for _, signal := range signals {
		for range signal {
		}
	}
}
