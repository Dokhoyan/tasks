package main

import "sort"

type ChangeKind string

const (
	Added   ChangeKind = "added"
	Removed ChangeKind = "removed"
	Updated ChangeKind = "updated"
)

type Change struct {
	Kind     ChangeKind
	Key      string
	OldValue string
	NewValue string
}

func Diff(old, current map[string]string) []Change {
	res := make([]Change, 0)

	for key, oldValue := range old {
		newValue, ok := current[key]

		if !ok {
			res = append(res, Change{
				Kind:     Removed,
				Key:      key,
				OldValue: oldValue,
			})
		} else if oldValue != newValue {
			res = append(res, Change{
				Kind:     Updated,
				Key:      key,
				OldValue: oldValue,
				NewValue: newValue,
			})
		}
	}

	for key, newValue := range current {
		if _, ok := old[key]; !ok {
			res = append(res, Change{
				Kind:     Added,
				Key:      key,
				NewValue: newValue,
			})
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Key < res[j].Key
	})

	return res
}
