package main

import "fmt"

// Нужно считывать сообщения из in до закрытия канала отправителем и разбить их на разделы:
// End = true завершает текущий раздел, его Value игнорируется. Нужно сохранить пустые разделы;
// незавершенный раздел при закрытии дает nil и ошибку. Если сообщений нет, разделов тоже нет;
// слайсы результата не разделяют элементы.
type SectionItem struct {
	Value int
	End   bool
}

func ReadSections(in <-chan SectionItem) ([][]int, error) {
	res := [][]int{}

	r := []int{}
	for s := range in {
		if s.End == false {
			r = append(r, s.Value)
		} else {
			res = append(res, r)
			r = []int{}
		}
	}

	if len(r) > 0 {
		return nil, fmt.Errorf("ReadSections: %v", r)
	}
	return res, nil
}
