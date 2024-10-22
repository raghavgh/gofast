package lfu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLFU(t *testing.T) {
	cases := []struct {
		name        string
		cap         int
		ops         []func(lfu *LFU)
		expectedLen int
		tests       []struct {
			key      string
			expected interface{}
		}
	}{
		{
			name: "Test Case 1",
			cap:  3,
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
			cap:  3,
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
		{
			name: "Test Case 3",
			cap:  3,
			ops: []func(lfu *LFU){
				func(lfu *LFU) { lfu.Put("1", 1) },
				func(lfu *LFU) { lfu.Get("1") },
				func(lfu *LFU) { lfu.Get("1") },
				func(lfu *LFU) { lfu.Put("2", 2) },
				func(lfu *LFU) { lfu.Put("3", 3) },
				func(lfu *LFU) { lfu.Put("4", 4) },
			},
			expectedLen: 3,
			tests: []struct {
				key      string
				expected interface{}
			}{
				{"1", 1},
				{"2", nil},
				{"3", 3},
				{"4", 4},
			},
		},
		{
			name: "Test Case 4",
			cap:  3,
			ops: []func(lfu *LFU){
				func(lfu *LFU) { lfu.Put("A", "Apple") },
				func(lfu *LFU) { lfu.Put("B", "Banana") },
				func(lfu *LFU) { lfu.Put("C", "Cantaloupe") },
				func(lfu *LFU) { lfu.Get("A") },
				func(lfu *LFU) { lfu.Get("A") },
				func(lfu *LFU) { lfu.Get("B") },
				func(lfu *LFU) { lfu.Put("D", "Dragonfruit") },
			},
			expectedLen: 3,
			tests: []struct {
				key      string
				expected interface{}
			}{
				{"A", "Apple"},
				{"B", "Banana"},
				{"C", nil},
				{"D", "Dragonfruit"},
			},
		},
		{
			name: "Test Case 5",
			cap:  3,
			ops: []func(lfu *LFU){
				func(lfu *LFU) { lfu.Put("1", 100) },
				func(lfu *LFU) { lfu.Remove("1") },
			},
			expectedLen: 0,
			tests: []struct {
				key      string
				expected interface{}
			}{
				{"1", nil},
			},
		},
		{
			name: "Test Case 6: Update existing keys",
			ops: []func(lfu *LFU){
				func(lfu *LFU) { lfu.Put("1", 1) },
				func(lfu *LFU) { lfu.Put("2", 2) },
				func(lfu *LFU) { lfu.Put("1", 10) },
				func(lfu *LFU) { lfu.Put("3", 3) },
				func(lfu *LFU) { lfu.Put("4", 4) },
			},
			expectedLen: 3,
			cap:         3,
			tests: []struct {
				key      string
				expected interface{}
			}{
				{"1", 10},
				{"2", nil},
				{"3", 3},
				{"4", 4},
			},
		},
		{
			name: "Test Case 7: Keys have different frequencies",
			ops: []func(lfu *LFU){
				func(lfu *LFU) { lfu.Put("1", 1) },
				func(lfu *LFU) { lfu.Get("1") },
				func(lfu *LFU) { lfu.Get("1") },
				func(lfu *LFU) { lfu.Put("2", 2) },
				func(lfu *LFU) { lfu.Get("2") },
				func(lfu *LFU) { lfu.Put("3", 3) },
				func(lfu *LFU) { lfu.Put("4", 4) },
			},
			expectedLen: 3,
			cap:         3,
			tests: []struct {
				key      string
				expected interface{}
			}{
				{"1", 1},
				{"2", 2},
				{"3", nil},
				{"4", 4},
			},
		},
		{
			name: "Test Case 8: LFU with capacity 1",
			ops: []func(lfu *LFU){
				func(lfu *LFU) { lfu.Put("1", 1) },
				func(lfu *LFU) { lfu.Put("2", 2) },
				func(lfu *LFU) { lfu.Put("3", 3) },
			},
			cap:         1,
			expectedLen: 1,
			tests: []struct {
				key      string
				expected interface{}
			}{
				{"1", nil},
				{"2", nil},
				{"3", 3},
			},
		},
		{
			name: "Test Case 9: All items appear equally",
			cap:  3,
			ops: []func(lfu *LFU){
				func(lfu *LFU) { lfu.Put("1", 1) },
				func(lfu *LFU) { lfu.Get("1") },
				func(lfu *LFU) { lfu.Put("2", 2) },
				func(lfu *LFU) { lfu.Get("2") },
				func(lfu *LFU) { lfu.Put("3", 3) },
				func(lfu *LFU) { lfu.Put("4", 4) },
				func(lfu *LFU) { lfu.Get("3") },
			},
			expectedLen: 3,
			tests: []struct {
				key      string
				expected interface{}
			}{
				{"1", 1},
				{"2", 2},
				{"3", nil},
				{"4", 4},
			},
		},
		{
			name: "Test Case 10: LFU with capacity 2",
			cap:  2,
			ops: []func(lfu *LFU){
				func(lfu *LFU) { lfu.Put("1", 1) },
				func(lfu *LFU) { lfu.Put("2", 2) },
				func(lfu *LFU) { lfu.Put("3", 3) },
			},
			expectedLen: 2,
			tests: []struct {
				key      string
				expected interface{}
			}{
				{"1", nil},
				{"2", 2},
				{"3", 3},
			},
		},
		{
			name: "Test Case 11: LFU with capacity 0",
			cap:  0,
			ops: []func(lfu *LFU){
				func(lfu *LFU) { lfu.Put("1", 1) },
			},
			expectedLen: 0,
			tests: []struct {
				key      string
				expected interface{}
			}{
				{"1", nil},
			},
		},
		{
			name: "Test Case 12: Set same key multiple times and different capacity",
			cap:  5,
			ops: []func(lfu *LFU){
				func(lfu *LFU) { lfu.Put("1", 1) },
				func(lfu *LFU) { lfu.Put("1", 10) },
				func(lfu *LFU) { lfu.Put("1", 100) },
				func(lfu *LFU) { lfu.Put("1", 1000) },
				func(lfu *LFU) { lfu.Get("1") },
				func(lfu *LFU) { lfu.Put("2", 2) },
			},
			expectedLen: 2,
			tests: []struct {
				key      string
				expected interface{}
			}{
				{"1", 1000},
				{"2", 2},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			lfu := NewLFU(tt.cap)
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
