package pokecache

import (
	"sync"
	"time"
	"errors"

)
type cacheEntry struct {
	createdAt  time.Time
	val []byte
}

type Cache struct {
	mu sync.Mutex
	entries map[string]cacheEntry
	interval time.Duration
	
}

func NewCache(interval int) *Cache {
	return &Cache{
		entries : make(map[string]cacheEntry),
		interval: time.Duration(interval),
	}

}

func (c *Cache) Add (key string, val []byte) error {

	if len(key) <= 0 {
		return errors.New("empty key is not allowed")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry {
		createdAt: time.Now(),
		val : val,
	}
	c.entries[key] = entry

	return nil

}

func (c *Cache) Get (key string) ([]byte , bool) {

	if len(key) <= 0 {
		return nil, false
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	
	result, exist := c.entries[key]

	if !exist {
		return nil, false
	}

	return result.val, exist



}
