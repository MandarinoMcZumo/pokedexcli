package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	interval time.Duration
	mu       *sync.Mutex
	data     map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(cacheDuration time.Duration) Cache {
	c := Cache{interval: cacheDuration, data: map[string]cacheEntry{}, mu: &sync.Mutex{}}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) error {
	newEntry := cacheEntry{createdAt: time.Now(), val: val}
	c.mu.Lock()
	c.data[key] = newEntry
	c.mu.Unlock()
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheHit, ok := c.data[key]
	if !ok {
		fmt.Println("cache hit not found for key: " + key)
		return []byte{}, ok
	}
	return cacheHit.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for t := range ticker.C {
		// fmt.Println("checking cache...")
		c.mu.Lock()
		for key, val := range c.data {
			maxTime := val.createdAt.Add(c.interval)
			if t.After(maxTime) {
				// fmt.Println("deleted cache key: " + key)
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}
