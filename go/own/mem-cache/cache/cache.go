package cache

import (
	"context"
	"errors"
	"sync"
	"time"
)

type CacheKey string
type CacheValue interface {
	CloneToCacheValue() CacheValue
}

var ErrNotFound = errors.New("not found")
var ErrExpired = errors.New("expired")

// В кэше используется мапа, операции с ней не потокобезопасны,
// поэтому операции критической секции синхронизируем спомощью мьютекса
type Cache struct {
	data map[CacheKey]*CacheItem
	ttl  time.Duration
	mu   sync.RWMutex
}

type CacheItem struct {
	value    CacheValue
	expireAt time.Time
}

func NewCache(ctx context.Context, ttl time.Duration) *Cache {
	cache := &Cache{
		data: make(map[CacheKey]*CacheItem),
		ttl:  ttl,
		mu:   sync.RWMutex{},
	}
	go cache.cleanup(ctx)
	return cache
}

func (c *Cache) Has(key CacheKey) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.data[key]
	return ok
}

func (c *Cache) Get(key CacheKey) (CacheValue, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.data[key]
	if !ok {
		return nil, ErrNotFound
	}
	if item.expireAt.Before(time.Now()) {
		return nil, ErrExpired
	}
	return item.value, nil
}

func (c *Cache) Set(key CacheKey, value CacheValue) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	expireAt := time.Now().Add(c.ttl)
	copied := value.CloneToCacheValue()
	c.data[key] = &CacheItem{copied, expireAt}
	return nil
}

func (c *Cache) cleanup(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// Find keys to delete from cache
			c.mu.RLock()
			delete_keys := []CacheKey{}
			for key, item := range c.data {
				if item.expireAt.Before(time.Now()) {
					delete_keys = append(delete_keys, key)
				}
			}
			c.mu.RUnlock()

			// Delete items by keys
			c.mu.Lock()
			for _, key := range delete_keys {
				delete(c.data, key)
			}
			c.mu.Unlock()
		case <-ctx.Done():
			return
		}
	}
}
