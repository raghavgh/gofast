package gofast

type Cache interface {
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
}
