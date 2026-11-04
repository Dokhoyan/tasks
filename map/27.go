package main

import "sort"

type Row struct {
	ID    int
	Name  string
	Score int
}

func InnerJoin(left map[int]string, right map[int]int) []Row {
	res := make([]Row, 0)

	for k, v := range left {
		if _, ok := right[k]; !ok {
			continue
		}
		res = append(res, Row{k, v, right[k]})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].ID < res[j].ID
	})

	return res
}
