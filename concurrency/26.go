package main

//Нужно считать строки из in до закрытия канала отправителем и вернуть те, что получены после
//первого marker. Само первое совпадение нужно исключить, последующие marker сохранить.
//Если совпадений нет, нужно вернуть пустой результат.

func AfterMarker(in <-chan string, marker string) []string {
	markerBeen := false
	res := make([]string, 0)

	for v := range in {
		if markerBeen {
			res = append(res, v)
		} else {
			if v == marker {
				markerBeen = true
			}
		}
	}
	return res
}
