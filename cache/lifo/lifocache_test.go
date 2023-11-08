package lifo

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLifo(t *testing.T) {
	t.Run("limit 1", func(t *testing.T) {
		l := NewLifo(1)
		assert.NotNil(t, l)

		t.Run("put 1", func(t *testing.T) {
			l.Put("1", 1)
			val, contains := l.Get("1")
			assert.True(t, contains)
			assert.Equal(t, 1, val.(int))

			assert.Equal(t, 1, l.Len())
			assert.True(t, l.Contains("1"))
		})

		t.Run("put 2", func(t *testing.T) {
			l.Put("2", 2)
			val, contains := l.Get("2")
			assert.True(t, contains)
			assert.Equal(t, 2, val.(int))

			_, contains = l.Get("1")
			assert.False(t, contains) // "1" should be evicted because of limit 1.

			assert.Equal(t, 1, l.Len())
			assert.True(t, l.Contains("2"))
			assert.False(t, l.Contains("1"))
		})

		t.Run("remove 2", func(t *testing.T) {
			l.Remove("2")
			_, contains := l.Get("2")
			assert.False(t, contains)

			assert.Equal(t, 0, l.Len())
			assert.False(t, l.Contains("2"))
		})

		t.Run("clear", func(t *testing.T) {
			l.Clear()
			assert.Equal(t, 0, l.Len())
		})
	})
	t.Run("limit 2", func(t *testing.T) {
		l := NewLifo(2)
		assert.NotNil(t, l)

		t.Run("Put string", func(t *testing.T) {
			l.Put("1", "s")
			val, contains := l.Get("1")
			assert.True(t, contains)
			assert.Equal(t, "s", val.(string))
		})

		t.Run("Put integer", func(t *testing.T) {
			l.Put("2", 1)
			val, contains := l.Get("2")
			assert.True(t, contains)
			assert.Equal(t, 1, val.(int))

			l.Put("3", 1)
			val, contains = l.Get("3")
			assert.True(t, contains)
			assert.Equal(t, 1, val.(int))

			_, contains = l.Get("2") // It should be evicted because the limit is 2
			assert.False(t, contains)
		})

		t.Run("Update existing key value", func(t *testing.T) {
			l.Put("2", 2)
			val, contains := l.Get("2")
			assert.True(t, contains)
			assert.Equal(t, 2, val.(int))
		})

		t.Run("Remove non-existing key", func(t *testing.T) {
			l.Remove("1") // "1" does not exist in the cache, so this operation should not panic
		})

		t.Run("Concurrency safety", func(t *testing.T) {
			wg := sync.WaitGroup{}
			for i := 0; i < 100; i++ {
				wg.Add(2)
				go func() {
					l.Get("2")
					wg.Done()
				}()
				go func() {
					l.Put("3", 3)
					wg.Done()
				}()
			}
			wg.Wait()

			val, contains := l.Get("3")
			assert.True(t, contains)
			assert.Equal(t, 3, val.(int))
		})
	})
}
