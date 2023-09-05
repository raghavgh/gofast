# gofast
This module prvide some easy to use inmemory cache algorithms implementations in golang like LRU, to use it you can simply import and use it. Follow given below instruction to use.

# How To use them?
1. Import it using this command
   go get github.com/raghavgh/gofast

2. Code eg. you can use given cache interface like this
   type CacheRepo struct {
	    MetaCache gofast.Cache
  }

  func NewCacheRepo() *CacheRepo {
	  return &CacheRepo{
      // 1 is algortithm type
	  	MetaCache: gofast.NewCache(1000, 1),
	  }
  }
3. Available Operations currently
  // Get returns the value (if any) and a boolean representing whether the value was found or not
	Get(key string) (any, bool)
	// Put adds a value to the cache
	Put(key string, val any)
	// Remove removes a value from the cache
	Remove(key string)
	// Len returns the number of elements of the cache
	Len() int
	// Clear clears the cache
	Clear()
	// Contains returns true if the cache contains the given key
	Contains(key string) bool

4. Currently Available cache algorithms
  a. LRU
  LFU and more number of algorithms will be available soon in new versions wait for it.

6. Given Algorithm types
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
	// SLRU is the segmented least recently used cache algorithm.
	SLRU
	// LIFO is the last in first out cache algorithm.
	LIFO

# Want to contribute?
Contributions will be welcomed very soon.


