package main

//Нужно отправить value в out и дождаться закрытия processed, подтверждающего обработку.
//Получатель работает параллельно и закроет processed без отправки значений после обработки;
//out открыт и может иметь буфер. Переданные каналы нельзя закрывать.

func SendAndWait(value int, out chan<- int, processed <-chan struct{}) {
	out <- value
	<-processed
}
