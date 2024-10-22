package lfu

import (
	"sync"

	"github.com/raghavgh/gofast/internal/ds/linkedlist"
)

type LFU struct {
	items         map[string]*entry
	freqToListMap map[int]*linkedlist.LinkedList
	minFreq       int
	mu            *sync.RWMutex
	limit         int
}

type entry struct {
	key   string
	value any
	freq  int
	node  *linkedlist.Node
}

func NewLFU(limit int) *LFU {
	return &LFU{
		items:         make(map[string]*entry),
		freqToListMap: make(map[int]*linkedlist.LinkedList),
		minFreq:       1,
		mu:            &sync.RWMutex{},
		limit:         limit,
	}
}

func (l *LFU) Get(key string) (any, bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()

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

	if l.limit == 0 {
		return
	}

	if element, ok := l.items[key]; ok {
		element.value = val
		l.updateFrequency(element)
		return
	}
	if len(l.items) >= l.limit {
		list := l.freqToListMap[l.minFreq]
		firstItemKey := list.Head.Val.(*entry).key
		list.Remove(list.Head)
		delete(l.items, firstItemKey)
	}

	dataEntry := &entry{
		key:   key,
		value: val,
		freq:  1,
	}
	l.addEntryInFreqList(dataEntry, dataEntry.freq)
	l.items[key] = dataEntry
	l.minFreq = 1
}

// Remove deletes a specific key-value pair from the cache.
func (l *LFU) Remove(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if element, ok := l.items[key]; ok {
		list := l.freqToListMap[element.freq]
		list.Remove(element.node)
		if list.Len() == 0 {
			delete(l.freqToListMap, element.freq)
		}
		delete(l.items, key)
	}
}

// Len returns the number of items in the cache.
func (l *LFU) Len() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return len(l.items)
}

// Clear removes all items from the cache.
func (l *LFU) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.items = make(map[string]*entry)
	l.freqToListMap = make(map[int]*linkedlist.LinkedList)
	l.minFreq = 1
}

// Contains returns true if the cache contains the given key
func (l *LFU) Contains(key string) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	_, ok := l.items[key]
	return ok
}

// updateFrequency updates the frequency of the given entry
func (l *LFU) updateFrequency(val *entry) {
	list := l.freqToListMap[val.freq]
	list.Remove(val.node)
	if list.Len() == 0 {
		delete(l.freqToListMap, val.freq)
		if l.minFreq == val.freq {
			l.minFreq++
		}
	}
	val.freq++
	l.addEntryInFreqList(val, val.freq)
}

// addEntryInFreqList updates the list for the given frequency
func (l *LFU) addEntryInFreqList(val *entry, freq int) {
	if _, ok := l.freqToListMap[freq]; !ok {
		l.freqToListMap[freq] = linkedlist.New()
	}
	node := l.freqToListMap[freq].PushBack(val)
	val.node = node
}
