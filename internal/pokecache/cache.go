package pokecache

import (
	"errors"
	"time"
)

type Cache struct {
	store map[string]cacheEntry
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache() Cache {
	return Cache{
		store: make(map[string]cacheEntry),
	}
}

// add to Cache
func (cache *Cache) AddDataToCache(url string, data []byte) {
	entry := cacheEntry{}
	entry.createdAt = time.Now()
	entry.val = data

	cache.store[url] = entry
}

// get from Cache
func (cache *Cache) GetDataFromCache(url string) ([]byte, error) {
	elem, ok := cache.store[url]
	if !ok {
		return nil, errors.New("No entry found")
	}
	return elem.val, nil
}

// Read loop
