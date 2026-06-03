package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		cachedData: make(map[string]cacheEntry),
		interval:   interval,
	}
	go newCache.reapLoop()
	return newCache
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

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		cutoff := time.Now().Add(-c.interval)
		for entry, data := range c.cachedData {
			if data.createdAt.Before(cutoff) {
				delete(c.cachedData, entry)
			}
		}
		c.mu.Unlock()
	}
}
