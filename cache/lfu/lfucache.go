package lfu

import (
	"sync"

	"github.com/raghavgh/gofast/ds/linkedlist"
)

type LFU struct {
	items   map[string]*entry
	freq    map[int]*linkedlist.LinkedList
	minFreq int
	mu      *sync.RWMutex
	limit   int
}

type entry struct {
	key   string
	value any
	freq  int
	node  *linkedlist.Node
}

// NewLFU returns a new LFU cache with given limit.
func NewLFU(limit int) *LFU {
	return &LFU{
		items:   make(map[string]*entry),
		freq:    make(map[int]*linkedlist.LinkedList),
		minFreq: 1,
		mu:      &sync.RWMutex{},
		limit:   limit,
	}
}

// Get retrieves a value from the cache for a specific key.
func (l *LFU) Get(key string) (any, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if element, ok := l.items[key]; ok {
		l.updateFrequency(element)
		return element.value, true
	}
	return nil, false
}

// Put adds a new key-value pair to the cache.
func (l *LFU) Put(key string, val any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if element, ok := l.items[key]; ok {
		element.value = val
		l.updateFrequency(element)
		return
	}
	if len(l.items) >= l.limit {
		list := l.freq[l.minFreq]
		firstItemKey := list.Head.Val.(*entry).key
		list.Remove(list.Head)
		delete(l.items, firstItemKey)
	}

	dataEntry := &entry{
		key:   key,
		value: val,
		freq:  1,
	}
	if list, ok := l.freq[dataEntry.freq]; ok {
		node := list.PushBack(dataEntry)
		dataEntry.node = node
	} else {
		list := linkedlist.New()
		node := list.PushBack(dataEntry)
		dataEntry.node = node
		l.freq[dataEntry.freq] = list
	}
	l.items[key] = dataEntry
	l.minFreq = 1
}

// Remove deletes a specific key-value pair from the cache.
func (l *LFU) Remove(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if element, ok := l.items[key]; ok {
		list := l.freq[element.freq]
		list.Remove(element.node)
		delete(l.items, key)
	}
}

// Len returns the number of elements of the cache.
func (l *LFU) Len() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return len(l.items)
}

// Clear clears the cache.
func (l *LFU) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.items = make(map[string]*entry)
	l.freq = make(map[int]*linkedlist.LinkedList)
	l.minFreq = 1
}

// Contains returns true if the cache contains the given key.
func (l *LFU) Contains(key string) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	_, ok := l.items[key]
	return ok
}

// updateFrequency u
func (l *LFU) updateFrequency(val *entry) {
	initNewList := func(val *entry) {
		tempList := linkedlist.New()
		node := tempList.PushBack(val)
		val.node = node
		l.freq[val.freq] = tempList
	}
	if list, ok := l.freq[val.freq]; ok {
		list.Remove(val.node)
		if list.Len() == 0 {
			delete(l.freq, val.freq)
			if val.freq == l.minFreq {
				l.minFreq++
			}
		}
	} else {
		initNewList(val)
	}
	val.freq++
	if newList, ok := l.freq[val.freq]; ok {
		node := newList.PushBack(val)
		val.node = node
	} else {
		initNewList(val)
	}

}
