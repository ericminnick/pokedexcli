package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	content 	map[string]cacheEntry
	mu 			*sync.Mutex
}

type cacheEntry struct {
	createdAt 	time.Time
	val 		[]byte
}


func NewCache(interval time.Duration) Cache {
	c := Cache{
		content: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.content[key] = cacheEntry{
		createdAt: 	time.Now().UTC(),
		val: 		value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.content[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.content {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.content, k)
		}
	}
}
