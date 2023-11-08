package fifo

import (
	"log"
	"sync"

	"github.com/raghavgh/gofast/ds/queue"
)

type Fifo struct {
	items             map[string]*entry
	queueEvictionList *queue.List
	limit             int
	mu                *sync.RWMutex
}

// entry is used to hold a value in the eviction list.
type entry struct {
	key   string
	value any
}

// NewFifo returns a new fifo cache with the given limit.
func NewFifo(limit int) *Fifo {
	if limit <= 0 {
		log.Fatal("Limit should be greater than 0")
	}
	return &Fifo{
		items:             make(map[string]*entry),
		queueEvictionList: queue.NewQueueList(true),
		limit:             limit,
		mu:                &sync.RWMutex{},
	}
}

// Get retrieves a value from the cache for a specific key.
func (f *Fifo) Get(key string) (any, bool) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	if element, ok := f.items[key]; ok {
		return element.value, true
	}
	return nil, false
}

// Put adds a new key-value pair to the cache.
func (f *Fifo) Put(key string, val any) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if element, ok := f.items[key]; ok {
		element.value = val
		return
	}
	if len(f.items) >= f.limit {
		element := f.queueEvictionList.Front()
		f.queueEvictionList.Pop()
		delete(f.items, element.(*entry).key)
	}
	entryVal := &entry{key: key, value: val}
	f.queueEvictionList.Push(entryVal)
	f.items[key] = entryVal
}

// Remove deletes a specific key-value pair from the cache.
func (f *Fifo) Remove(key string) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if element, ok := f.items[key]; ok {
		delete(f.items, key)
		f.queueEvictionList.Remove(element)
	}
}

// Len returns the number of items in the cache.
func (f *Fifo) Len() int {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return len(f.items)
}

// Clear removes all items from the cache.
func (f *Fifo) Clear() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.items = make(map[string]*entry)
	f.queueEvictionList = queue.NewQueueList(true)
}

// Contains returns true if the cache contains the given key
func (f *Fifo) Contains(key string) bool {
	f.mu.RLock()
	defer f.mu.RUnlock()
	_, ok := f.items[key]
	return ok
}
