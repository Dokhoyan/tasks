package main

import "strings"

func Join(parts []string, sep string) string {
	if len(parts) == 0 {
		return ""
	}
	res := strings.Builder{}
	for _, part := range parts[:len(parts)-1] {
		res.WriteString(part)
		res.WriteString(sep)
	}
	res.WriteString(parts[len(parts)-1])
	return res.String()
}
