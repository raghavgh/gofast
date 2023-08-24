package lru

import (
	"sync"

	"github.com/raghavgh/gofast/ds/linkedlist"
)

// LRU represents a thread-safe, least recently used cache.
type LRU struct {
	items    map[string]*linkedlist.Node
	eviction *linkedlist.LinkedList
	limit    int
	mu       *sync.RWMutex
}

// entry is used to hold a value in the eviction list.
// we are keeping entry as value to make sure that we can access key in O(1) time
// when we want to remove an entry from the map.
type entry struct {
	key   string
	value any
}

// Get retrieves a value from the cache for a specific key.
func (l *LRU) Get(key string) (any, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if node, ok := l.items[key]; ok {
		l.eviction.MoveToFront(node)
		return node.Val.(*entry).value, true
	}
	return nil, false
}

// Put adds a new key-value pair to the cache.
func (l *LRU) Put(key string, val any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.eviction.Len() >= l.limit {
		delete(l.items, l.eviction.Tail.Val.(*entry).key)
		l.eviction.Remove(l.eviction.Tail)
	}
	l.eviction.PushFront(&entry{key: key, value: val})
	l.items[key] = l.eviction.Head
}

// Remove deletes a specific key-value pair from the cache.
func (l *LRU) Remove(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if node, ok := l.items[key]; ok {
		delete(l.items, key)
		l.eviction.Remove(node)
	}
}

// Len returns the number of items in the cache.
func (l *LRU) Len() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.eviction.Len()
}

// Clear removes all items from the cache.
func (l *LRU) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Reset the eviction list and the map.
	// The old items will be garbage collected.
	l.items = make(map[string]*linkedlist.Node)
	l.eviction = linkedlist.New()
}

// Contains checks if a key is present in the cache.
func (l *LRU) Contains(key string) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	_, ok := l.items[key]
	return ok
}

// NewLRU creates a new LRU cache with the maximum size based on configuration.
func NewLRU(limit int) *LRU {
	return &LRU{
		items:    make(map[string]*linkedlist.Node, limit),
		eviction: linkedlist.New(),
		limit:    limit,
		mu:       &sync.RWMutex{},
	}
}
