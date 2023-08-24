package lru

import (
	"strconv"
	"testing"
)

func TestCache(t *testing.T) {
	cache := NewLRU(1000)

	t.Run("Put and Get", func(t *testing.T) {
		key := "key1"
		value := "value1"
		cache.Put(key, value)

		v, ok := cache.Get(key)
		if !ok || v != value {
			t.Errorf("Expected value %v, got %v", value, v)
		}
	})

	t.Run("Remove", func(t *testing.T) {
		key := "key2"
		value := "value2"
		cache.Put(key, value)

		cache.Remove(key)
		_, ok := cache.Get(key)
		if ok {
			t.Errorf("Expected key %s to be removed", key)
		}
	})

	t.Run("Clear", func(t *testing.T) {
		cache.Put("key", "value")

		cache.Clear()
		if cache.Len() != 0 {
			t.Errorf("Expected cache to be cleared")
		}
	})

	t.Run("Contains", func(t *testing.T) {
		key := "key3"
		value := "value3"
		cache.Put(key, value)

		if !cache.Contains(key) {
			t.Errorf("Expected cache to contain key %s", key)
		}
	})

	t.Run("Evict", func(t *testing.T) {
		cache.Clear()

		for i := 0; i < cache.limit; i++ {
			key := strconv.Itoa(i)
			cache.Put(key, "value")
		}

		// This should trigger an eviction of key "0"
		cache.Put("new", "value")

		_, ok := cache.Get("0")
		if ok {
			t.Errorf("Expected key '0' to be evicted")
		}
	})
}

func TestCache_Concurrent(t *testing.T) {
	cache := NewLRU(1000)

	t.Run("Concurrent Get and Put", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			go cache.Put("key", "value")
			go cache.Get("key")
		}
	})

	t.Run("Concurrent Remove", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			go cache.Put("key", "value")
			go cache.Remove("key")
		}
	})
}
