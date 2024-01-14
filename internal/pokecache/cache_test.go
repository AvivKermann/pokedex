package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.entries == nil {
		t.Error("cache is nil")
	}
}

func TestAddEntry(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cache.Add("key1", []byte("val1"))
	actual, ok := cache.Get("key1")

	if !ok {
		t.Error("key1 not found")
	}
	if string(actual) != "val1" {
		t.Error("value dont match")
	}

}

func TestReap(t *testing.T) {
	interval := time.Microsecond * 10
	cache := NewCache(interval)

	cache.Add("key1", []byte("val1"))
	time.Sleep(interval / 2)
	
	
	_, ok := cache.Get("key1")

	if ok {
		t.Error("should not  reap anything")
	}

}
