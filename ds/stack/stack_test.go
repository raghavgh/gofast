package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	tests := []struct {
		name         string
		elements     []int
		expectedSize int
		pop          bool
		expectedTop  int
		shouldErrTop bool
	}{
		{"Push one", []int{1}, 1, false, 1, false},
		{"Push two", []int{1, 2}, 2, false, 2, false},
		{"Push three", []int{1, 2, 3}, 3, false, 3, false},
		{"Pop one", []int{1}, 0, true, 0, true},
		{"Pop two", []int{1, 2}, 0, true, 0, true},
		{"Push and then Pop", []int{1, 2, 3}, 0, true, 0, true},
		{"Push multiple and then Pop", []int{1, 2, 3, 4, 5}, 0, true, 4, true},
		{"Pop until Empty", []int{1}, 0, true, 0, true},
		{"Keep Stack Empty", []int{}, 0, true, 0, true},
		{"Push and Pop Alternate", []int{1, 2, 3, 4, 5}, 0, true, 0, true},
		{"Examine Empty Stack", []int{}, 0, false, 0, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewStack()

			for _, el := range tc.elements {
				s.Push(el)
			}

			if tc.pop {
				for !s.Empty() {
					s.Pop()
				}
			}

			assert.Equal(t, tc.expectedSize, s.Size())

			if tc.shouldErrTop {
				assert.Panics(t, func() { s.Top() })
			} else {
				assert.Equal(t, tc.expectedTop, s.Top())
			}
		})
	}
}
