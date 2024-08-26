package cache

type CacheValue interface{}
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
	c.data[key] = value
	return nil
}

func (c *Cache) Get(key CacheKey) CacheValue {
	value, _ := c.data[key]
	return value
}
