package main

import (
	"fmt"
)

// Нужно считать ответы из in до закрытия канала отправителем и сопоставить с уникальными
// requestIDs. Нужно вернуть ответы в порядке запросов. Неизвестный ID, повтор или отсутствие
// ответа дают nil и ошибку, но не отменяют чтение до закрытия. Для пустого списка запросов
// корректен только канал без ответов.
type Reply struct {
	ID    string
	Value int
}

func MatchReplies(requestIDs []string, in <-chan Reply) ([]Reply, error) {
	positions := make(map[string]int, len(requestIDs))
	for i, id := range requestIDs {
		positions[id] = i
	}

	result := make([]Reply, len(requestIDs))
	seen := make([]bool, len(requestIDs))
	var firstErr error

	for reply := range in {
		index, ok := positions[reply.ID]
		if !ok {
			if firstErr == nil {
				firstErr = fmt.Errorf("unknown ID: %q", reply.ID)
			}
			continue
		}

		if seen[index] {
			if firstErr == nil {
				firstErr = fmt.Errorf("duplicate reply: %q", reply.ID)
			}
			continue
		}

		result[index] = reply
		seen[index] = true
	}

	if firstErr != nil {
		return nil, firstErr
	}

	for i, received := range seen {
		if !received {
			return nil, fmt.Errorf("missing reply: %q", requestIDs[i])
		}
	}

	return result, nil
}
