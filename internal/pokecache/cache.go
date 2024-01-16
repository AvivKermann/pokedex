package pokecache

import (
	"errors"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu      sync.Mutex
	entries map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c

}

func (c *Cache) Add(key string, val []byte) error {

	if len(key) <= 0 {
		return errors.New("empty key is not allowed")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
	c.entries[key] = entry

	return nil

}

func (c *Cache) Get(key string) ([]byte, bool) {

	if len(key) <= 0 {
		return nil, false
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exist := c.entries[key]

	if !exist {
		return nil, false
	}

	return entry.val, exist
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {

	deletionTime := time.Now().UTC().Add(-interval)
	for key, value := range c.entries {

		if value.createdAt.Before(deletionTime) {
			delete(c.entries, key)
		}
	}
}
