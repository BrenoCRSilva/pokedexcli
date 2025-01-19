package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	storage map[string]cacheEntry
	mutex   sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func newCache(interval time.Duration) map[string]cacheEntry {
	var mu sync.Mutex
	cache := Cache{
		storage: make(map[string]cacheEntry),
		mutex:   mu,
	}
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	createdAt := time.Now()
	entry := cacheEntry{
		createdAt: createdAt,
		val:       val,
	}
	c.storage[key] = entry
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if entry, ok := c.storage[key]; !ok {
		return nil, false
	} else {
		return entry.val, true
	}
}

func (c *Cache) reapLoop() {
}
