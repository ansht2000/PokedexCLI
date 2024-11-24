package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache 	map[string]cacheEntry
	mut 	*sync.Mutex
}

type cacheEntry struct {
	createdAt 	time.Time
	val 		[]byte
}

// Create a new cache
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mut: &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

// Add a value to the cache
func (c *Cache) Add(key string, value []byte) {
	newEntry := cacheEntry{
		createdAt: time.Now().UTC(),
		val: value,
	}

	c.mut.Lock()
	defer c.mut.Unlock()
	c.cache[key] = newEntry
}

// Get a value from the cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()
	val, exists := c.cache[key]
	return val.val, exists
}

// Start the loop to delete values from the cache
// during each interval tick
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

// Delete values from the cache that have existed
// longer than the specified interval
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mut.Lock()
	defer c.mut.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}