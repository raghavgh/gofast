package gofast

import (
	"github.com/raghavgh/gofast/internal/cache/fifo"
	"github.com/raghavgh/gofast/internal/cache/lifo"
	"github.com/raghavgh/gofast/internal/cache/lru"
)

type Algorithm int

const (
	// LRU is the least recently used cache algorithm.
	LRU Algorithm = iota
	// LFU is the least frequently used cache algorithm.
	LFU
	// FIFO is the first in first out cache algorithm.
	FIFO
	// MRU is the most recently used cache algorithm.
	MRU
	// RR is the random replacement cache algorithm.
	RR
	// ARC is the adaptive replacement cache algorithm.
	ARC
	// SLRU is the segmented least recently used cache algorithm.
	SLRU
	// LIFO is the last in first out cache algorithm.
	LIFO
	// TTL is the time to live cache algorithm.
	TTL
)

// NewCache returns a new cache with the given limit and algorithm.
func NewCache(limit int, algo Algorithm) Cache {
	switch algo {
	case LRU:
		return lru.NewLRU(limit)
	case FIFO:
		return fifo.NewFifo(limit)
	case LFU:
		return lifo.NewLifo(limit)
	case LIFO:
		return lifo.NewLifo(limit)
	default:
		return lru.NewLRU(limit)
	}
}
