package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		name             string
		elements         []any
		expectedSize     int
		remove           bool
		expectedFront    any
		expectedBack     any
		expectedFrontErr bool
		expectedBackErr  bool
	}{
		{"Push one", []any{1}, 1, false, 1, 1, false, false},
		{"Push two", []any{1, 2}, 2, false, 1, 2, false, false},
		{"Push three", []any{1, 2, 3}, 3, false, 1, 3, false, false},
		{"Push four", []any{1, 2, 3, 4}, 4, false, 1, 4, false, false},
		{"Push five", []any{1, 2, 3, 4, 5}, 5, false, 1, 5, false, false},
		{"Remove one", []any{1}, 0, true, nil, nil, true, true},
		{"Remove two", []any{1, 2}, 1, true, 2, 2, false, false},
		{"Push different types", []any{1, "2", 3.0, 4, "5"}, 5, false, 1, "5", false, false},
		{"Remove different types", []any{1, "2", 3.0, 4, "5"}, 4, true, "2", "5", false, false},
		{"Queue is empty initially", []any{}, 0, false, nil, nil, true, true},
		{"Keep queue empty", []any{1}, 0, true, nil, nil, true, true},
		{"Push and Remove alternately", []any{1, 2, 3, 4, 5}, 4, true, 2, 5, false, false},
		{"Push 10 elements to queue", []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, false, 1, 10, false, false},
		{"Push and Remove with dupes", []any{1, 1, 2, 2, 3, 3}, 5, true, 1, 3, false, false},
		{"Push and check size", []any{1, 2, 3}, 3, false, 1, 3, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueueList()

			for _, element := range tt.elements {
				q.Push(element)
			}

			if tt.remove {
				q.Pop()
			}

			assert.Equal(t, tt.expectedSize, q.Size())

			if !tt.expectedFrontErr {
				assert.Equal(t, tt.expectedFront, q.Front())
			} else {
				assert.Panics(t, func() { q.Front() })
			}

			if !tt.expectedBackErr {
				assert.Equal(t, tt.expectedBack, q.Back())
			} else {
				assert.Panics(t, func() { q.Back() })
			}
		})
	}
}
