package mru

import (
	"sync"

	"github.com/raghavgh/gofast/internal/ds/linkedlist"
)

/*
MRU represetns a thread-safe, most recently used cache.
When a new item gets added, it will be the very first item in the cache
When a item get updated, it will move to the first place of the cache
When cache reach it's limit, it will remove the first item
*/
type MRU struct {
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
// And move the item to the front
func (m *MRU) Get(key string) (any, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if node, ok := m.items[key]; ok {
		m.eviction.MoveToFront(node)
		return node.Val.(*entry).value, true
	}

	return nil, false
}

// Put add or update the key-value pair to the cache.
// And move the item to front
func (m *MRU) Put(key string, val any) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if element, ok := m.items[key]; ok {
		m.eviction.MoveToFront(element)
		element.Val.(*entry).value = val
		return
	}

	if m.eviction.Len() >= m.limit {
		delete(m.items, m.eviction.Head.Val.(*entry).key)
		m.eviction.Remove(m.eviction.Head)
	}

	m.eviction.PushFront(&entry{key: key, value: val})
	m.items[key] = m.eviction.Head
}

// Remove deletes a specific key-value pair from the cache.
func (m *MRU) Remove(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if node, ok := m.items[key]; ok {
		delete(m.items, key)
		m.eviction.Remove(node)
	}
}

// Len returns the number of items in the cache.
func (m *MRU) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.eviction.Len()
}

// Clear removes all items from the cache.
func (m *MRU) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.items = make(map[string]*linkedlist.Node)
	m.eviction = linkedlist.New()
}

// Contains checks if a key is present in the cache.
func (m *MRU) Contains(key string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.items[key]
	return ok
}

// NewMRU creates a new MRU cache with the maximum size based on configuration.
func NewMRU(limit int) *MRU {
	if limit <= 0 {
		panic("cache limit must be greater than 0")
	}

	return &MRU{
		items:    make(map[string]*linkedlist.Node, limit),
		eviction: linkedlist.New(),
		limit:    limit,
		mu:       &sync.RWMutex{},
	}
}
