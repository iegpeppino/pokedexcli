package pokecache

import (
	"sync"
	"time"
)

/*
I defer the mu unlocks for any strange case
where something unexpected happens so the cache
doesn't get locked indeterminately
*/

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mu:      &sync.Mutex{},
		entries: map[string]cacheEntry{},
	}
	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	// Add a new entry to the cache
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.entries[key]
	// If the entry exists, it returns the value and true
	if exists {
		return entry.val, exists
	}
	// If it doesn't it returns nil and false
	return nil, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	// Create ticker
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()

		for key, entry := range c.entries {
			t := time.Now()
			// If the entry has been in the cache for longer
			// than the interval, it is deleted
			if t.Sub(entry.createdAt) > interval {
				delete(c.entries, key)
			}

		}
		c.mu.Unlock()
	}
	// Periodically cleans the cache

}
