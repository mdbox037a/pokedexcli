package pokecache

import "time"

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		cachedData: make(map[string]cacheEntry),
		interval:   interval,
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cachedData[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) (data []byte, found bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if entry, ok := c.cachedData[key]; ok {
		return entry.val, true
	}
	return nil, false
}
