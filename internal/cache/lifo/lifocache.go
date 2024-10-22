package lifo

import (
	"log"
	"sync"

	"github.com/raghavgh/gofast/internal/ds/stack"
)

// Lifo represents a lifo cache.
type Lifo struct {
	items map[string]*entry
	stack *stack.Stack
	limit int
	mu    *sync.RWMutex
}

// entry is used to hold a value in the eviction list.
type entry struct {
	key   string
	value any
}

// NewLifo returns a new lifo cache with the given limit.
func NewLifo(limit int) *Lifo {
	if limit <= 0 {
		log.Fatal("Limit should be greater than 0")
	}
	return &Lifo{
		items: make(map[string]*entry),
		stack: stack.NewStack(true),
		limit: limit,
		mu:    &sync.RWMutex{},
	}
}

// Get retrieves a value from the cache for a specific key.
func (l *Lifo) Get(key string) (any, bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if element, ok := l.items[key]; ok {
		return element.value, true
	}
	return nil, false
}

// Put adds a new key-value pair to the cache.
func (l *Lifo) Put(key string, val any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if element, ok := l.items[key]; ok {
		element.value = val
		return
	}
	if l.limit <= l.stack.Size() {
		element := l.stack.Top()
		l.stack.Pop()
		delete(l.items, element.(*entry).key)
	}
	entryVal := &entry{
		key:   key,
		value: val,
	}
	l.stack.Push(entryVal)
	l.items[key] = entryVal
}

// Remove deletes a specific key-value pair from the cache.
func (l *Lifo) Remove(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if element, ok := l.items[key]; ok {
		l.stack.Remove(element)
		delete(l.items, key)
	}
}

// Len returns the number of items in the cache.
func (l *Lifo) Len() int {
	return l.stack.Size()
}

// Clear purges all stored items from the cache.
func (l *Lifo) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.stack.Clear()
	l.items = make(map[string]*entry)
}

// Contains returns true if the cache contains the given key.
func (l *Lifo) Contains(key string) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()

	_, ok := l.items[key]
	return ok
}
