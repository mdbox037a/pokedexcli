package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cachedData map[string]cacheEntry
	mu         sync.Mutex
	interval   time.Duration
}
