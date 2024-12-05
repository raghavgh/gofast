package mru

import (
	"strconv"
	"testing"
)

func TestCache(t *testing.T) {
	cache := NewMRU(1000)

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

		// This should trigger an eviction of key "9"
		cache.Put("new", "value")

		mostRecentAddedItemKey := strconv.Itoa(cache.limit - 1)
		_, ok := cache.Get(mostRecentAddedItemKey)
		if ok {
			t.Errorf("Expected key '%s' to be evicted", mostRecentAddedItemKey)
		}

		_, ok = cache.Get("0")
		if !ok {
			t.Errorf("Expected key '0' to still be present")
		}
	})

	t.Run("UpdateExistingKey", func(t *testing.T) {
		key := "key1"
		initialVal := "value1"
		updatedVal := "value2"
		cache.Put(key, initialVal)

		cache.Put(key, updatedVal) // Updating the value of the existing key

		v, ok := cache.Get(key)
		if !ok || v != updatedVal {
			t.Errorf("Expected updated value %v, got %v", updatedVal, v)
		}
	})

	t.Run("Evict", func(t *testing.T) {
		cache.Clear()

		for i := 0; i < cache.limit; i++ {
			key := strconv.Itoa(i)
			cache.Put(key, "value")
		}

		// This should replace the value of key "0" rather than evicting it
		cache.Put("0", "newValue")

		v, ok := cache.Get("0")
		if !ok || v != "newValue" {
			t.Errorf("Expected key '0' to be updated rather than evicted")
		}
	})
}

func TestCache_Concurrent(t *testing.T) {
	cache := NewMRU(1000)

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
