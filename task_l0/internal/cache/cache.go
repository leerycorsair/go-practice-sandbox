package cache

import (
	"sync"
)

type Cache struct {
	sync.RWMutex
	items map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{items: make(map[string]interface{})}
}

func (c *Cache) Add(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()
	c.items[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	item, found := c.items[key]
	if !found {
		return nil, false
	}
	return item, true
}

func (c *Cache) GetAll() []interface{} {
	var values []interface{}
	for _, v := range c.items {
		values = append(values, v)
	}
	return values
}
