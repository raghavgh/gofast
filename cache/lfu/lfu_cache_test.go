package lfu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLFU(t *testing.T) {
	cases := []struct {
		name        string
		ops         []func(lfu *LFU)
		expectedLen int
		tests       []struct {
			key      string
			expected interface{}
		}
	}{
		{
			name: "Test Case 1",
			ops: []func(lfu *LFU){
				func(lfu *LFU) { lfu.Put("1", 1) },
				func(lfu *LFU) { lfu.Put("2", 2) },
				func(lfu *LFU) { lfu.Put("3", 3) },
				func(lfu *LFU) { lfu.Get("1") },
				func(lfu *LFU) { lfu.Put("4", 4) },
				func(lfu *LFU) { lfu.Get("2") },
				func(lfu *LFU) { lfu.Put("5", 5) },
			},
			expectedLen: 3,
			tests: []struct {
				key      string
				expected interface{}
			}{
				{"1", 1},
				{"2", nil},
				{"3", nil},
				{"4", 4},
				{"5", 5},
			},
		},
		{
			name: "Test Case 2",
			ops: []func(lfu *LFU){
				func(lfu *LFU) { lfu.Clear() },
			},
			expectedLen: 0,
			tests: []struct {
				key      string
				expected interface{}
			}{
				{"1", nil},
				{"2", nil},
				{"3", nil},
				{"4", nil},
				{"5", nil},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			lfu := NewLFU(3)
			for _, op := range tt.ops {
				op(lfu)
			}
			assert.Equal(t, tt.expectedLen, lfu.Len())
			for _, test := range tt.tests {
				val, _ := lfu.Get(test.key)
				assert.Equal(t, test.expected, val)
			}
		})
	}
}
