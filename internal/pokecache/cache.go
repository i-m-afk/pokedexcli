package pokecache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	store map[string]cacheEntry
	mux   sync.Mutex
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) *Cache {
	c := Cache{
		store: make(map[string]cacheEntry),
		mux:   sync.Mutex{},
	}
	go c.reapLoop(duration)
	return &c
}

// add to Cache
func (cache *Cache) AddDataToCache(url string, data []byte) {
	entry := cacheEntry{}
	entry.createdAt = time.Now()
	entry.val = data

	cache.mux.Lock()
	cache.store[url] = entry
	cache.mux.Unlock()
}

// get from Cache
func (cache *Cache) GetDataFromCache(url string) ([]byte, error) {
	cache.mux.Lock()
	elem, ok := cache.store[url]
	cache.mux.Unlock()
	if !ok {
		return nil, errors.New("No entry found")
	}
	return elem.val, nil
}

// Reap loop
func (cache *Cache) reapLoop(duration time.Duration) {
	tick := time.NewTicker(duration)
	for range tick.C {
		cache.reap(duration)
	}
}

func (cache *Cache) reap(duration time.Duration) {
	cache.mux.Lock()
	for key, val := range cache.store {
		if time.Since(val.createdAt) >= duration {
			delete(cache.store, key)
		}
	}
	cache.mux.Unlock()
}

// cache invalidation
// func (cache *Cache)
