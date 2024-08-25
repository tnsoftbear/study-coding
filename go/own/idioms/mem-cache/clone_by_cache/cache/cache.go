package cache

type CacheValue interface {
	CloneToCacheValue() CacheValue
}
type CacheKey string

type Cache struct {
	data map[CacheKey]CacheValue
}

func NewCache() *Cache {
	cache := &Cache{
		data: make(map[CacheKey]CacheValue),
	}
	return cache
}

func (c *Cache) Set(key CacheKey, value CacheValue) error {
	clone := value.CloneToCacheValue()
	c.data[key] = clone
	return nil
}

func (c *Cache) Get(key CacheKey) CacheValue {
	value := c.data[key]
	return value
}
