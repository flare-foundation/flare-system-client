package utils

import (
	"sync"
)

type CacheBase[K comparable, V any] interface {
	Add(K, V)
	Get(K) (V, bool)
}

type Cache[K comparable, V any] interface {
	CacheBase[K, V]

	RemoveAccessed()
}

// Map object cache
type cache[K comparable, V any] struct {
	sync.RWMutex

	cacheMap map[K]V
	accessed []K
}

func NewCache[K comparable, V any]() Cache[K, V] {
	return &cache[K, V]{
		cacheMap: make(map[K]V),
		accessed: nil,
	}
}

func (c *cache[K, V]) Add(k K, v V) {
	c.cacheMap[k] = v
}

func (c *cache[K, V]) Get(k K) (V, bool) {
	c.RWMutex.Lock()
	v, ok := c.cacheMap[k]
	if ok {
		c.accessed = append(c.accessed, k)
	}
	c.RWMutex.Unlock()
	return v, ok
}

func (c *cache[K, V]) RemoveAccessed() {
	c.RWMutex.Lock()
	for _, k := range c.accessed {
		delete(c.cacheMap, k)
	}
	c.accessed = nil
	c.RWMutex.Unlock()
}
