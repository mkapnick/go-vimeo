package lrucache

import (
	"time"
)

type CacheItem struct {
	CreatedAt time.Time
	LastUsed  time.Time
	Key       string
	Value     []byte
	Size      int
}

// LRU cache
type Cache struct {
	Storage map[string]CacheItem
	MaxSize int
	Size    int
}

func New(size int) *Cache {
	// size of cache storage should not be greater
	// than 64MB
	if size > 64 {
		// default to 64 MB
		size = 64
	}

	return &Cache{
		MaxSize: size * 1000000,
		Size:    0,
	}
}

func (c *Cache) Push(key string, value []byte) (bool, error) {
	size := len(value)

	if c.Size+size > c.MaxSize {
		// pop the least recently used from the cache
	}

	// otherwise everything

	return false, nil
}
