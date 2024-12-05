# gofast

gofast is a Go module that provides easy-to-use in-memory cache algorithms like LRU (Least Recently Used). It allows you to import and use these algorithms simply and directly in your Go projects.

## Key Features:
1. <B>High Speed</B> : Optimized for low-latency operations, making it one of the fastest in-memory caching solutions available.
2. <B>Thread-Safe</B> : Designed with concurrency in mind, allowing multiple goroutines to safely access and modify the cache.
3. <B>Simple API</B> : Offers a straightforward and easy-to-use API for basic caching operations like Get, Set, and Remove.
4. <B>Customizable Eviction Policy</B>: Supports options like LRU (Least Recently Used) for memory management when cache size reaches its limit.
5. <B>Lightweight</B> : Minimal external dependencies, ensuring a lean and performant implementation.
6. <B>Custom Optimized Data Structures</B> : Leverages its own highly efficient in-built data structures, specifically tailored for time efficiency and reduced overhead, giving it a competitive edge in speed and performance.

## Installation

Install gofast by simply using the `go get` command:

```bash
$ go get github.com/raghavgh/gofast
```

## Usage
### Importing the module
First, import gofast like this:

```go
import "github.com/raghavgh/gofast"
```
Here is an example of how you can use the module:
```go
type CacheRepo struct {
    MetaCache gofast.Cache
}

func NewCacheRepo() *CacheRepo {
    return &CacheRepo{
        // gofast.LRU is the algorithm type for LRU.
	// 1000 is limit of your cache.
        MetaCache: gofast.NewCache(1000, gofast.LRU),
    }
}
```
### Available Functions
```go
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
```
**All above funtions are thread safe

### Available Cache Algorithms
Currently, the following cache algorithms are available:

1. LRU (Least Recently Used)

2. LFU (Least Frequently Used)

3. LIFO (Last In, First Out)

4. FIFO (First In, First Out)

5. MRU (Most Recently Used) 

More algorithms will be available in future versions.

Supported algorithms can be specified with the following constants:

```go
gofast.LRU   // Least Recently Used algorithm
gofast.LFU   // Least Frequently Used algorithm
gofast.FIFO  // First In, First Out algorithm
gofast.MRU   // Most Recently Used algorithm
gofast.RR    // Random Replacement algorithm
gofast.SLRU  // Segmented Least Recently Used algorithm
gofast.LIFO  // Last In, First Out algorithm
```
## 🤝 Contributions Welcome!
We’re excited to have you contribute to the gofast library! Whether you’re fixing bugs 🐛, adding new features ✨, improving documentation 📚, or enhancing test coverage 🧪—we’d love your help!

#### How to Get Started:

1. 🍴 Fork the repository and create your feature branch (git checkout -b feature/AmazingFeature).
2. 🛠️ Make your changes and commit them (git commit -m 'Add some AmazingFeature').
3. ⏫ Push to the branch (git push origin feature/AmazingFeature).
4. 🔍 Open a pull request, and we’ll review it as soon as possible!

#### Looking for Ideas?

Check out our [issues page](https://github.com/raghavgh/gofast/issues) for a list of open tasks or feel free to suggest your own improvements! We’re especially interested in:

- 🧩 Implementing new cache algorithms (e.g., MRU). 
- 🔍 Adding more unit tests for better coverage. 
- ⚡ Improving performance with enhanced benchmarking.

We look forward to your valuable contributions and ideas 💡! Let’s make gofast even faster together! 🎉
