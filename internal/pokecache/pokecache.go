package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	store    map[string]cacheEntry
	interval time.Duration
	mux      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		store:    make(map[string]cacheEntry),
		interval: interval,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	createdAt := time.Now()
	entry := cacheEntry{
		createdAt: createdAt,
		val:       val,
	}
	c.store[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if entry, ok := c.store[key]; !ok {
		return nil, false
	} else {
		return entry.val, true
	}
}

func (c *Cache) reapLoop() {
	c.mux.Lock()
	defer c.mux.Unlock()
	time.Sleep(c.interval)
	now := time.Now()
	for k, v := range c.store {
		if now.Sub(v.createdAt) > 5*time.Second {
			delete(c.store, k)
		}
	}
}
