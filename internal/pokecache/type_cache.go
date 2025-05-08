package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry        map[string]cacheEntry
	mu           sync.Mutex
	reapInterval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		entry:        make(map[string]cacheEntry),
		reapInterval: interval,
	}
	go c.reapLoop()
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entry[key] = cacheEntry{
		time.Now(),
		val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	val, ok := c.entry[key]
	if !ok {
		c.mu.Unlock()
		return nil, false
	}
	c.mu.Unlock()
	return val.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.reapInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			for key, val := range c.entry {
				if time.Since(val.createdAt) > c.reapInterval {
					delete(c.entry, key)
				}
			}
			c.mu.Unlock()
		}
	}
}
