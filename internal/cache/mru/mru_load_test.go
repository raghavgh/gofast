package mru

import (
	"strconv"
	"testing"
)

// Initialize your MRU cache
var cache = NewMRU(1000)

func BenchmarkPut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache.Put("Key"+strconv.Itoa(i), i)
	}
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache.Get("Key" + strconv.Itoa(i))
	}
}

func BenchmarkRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache.Remove("Key" + strconv.Itoa(i))
	}
}

func BenchmarkContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache.Contains("Key" + strconv.Itoa(i))
	}
}

func BenchmarkClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache.Clear()
	}
}

func BenchmarkLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache.Len()
	}
}
