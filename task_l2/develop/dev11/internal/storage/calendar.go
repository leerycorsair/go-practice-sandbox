package storage

import (
	"develop/dev11/internal/models"
	"sync"
)

type CalendarStorage struct {
	m    *sync.RWMutex
	data map[int]models.EventT
}

func NewCalendarStorage(m *sync.RWMutex) *CalendarStorage {
	return &CalendarStorage{
		m:    m,
		data: map[int]models.EventT{},
	}
}

func (c *CalendarStorage) Get(key int) (models.EventT, bool) {
	c.m.RLock()
	defer c.m.RUnlock()
	entity, ok := c.data[key]
	return entity, ok
}

func (c *CalendarStorage) Set(key int, event models.EventT) {
	c.m.Lock()
	defer c.m.Unlock()
	c.data[key] = event
}

func (c *CalendarStorage) Len() int {
	c.m.RLock()
	defer c.m.RUnlock()
	return len(c.data)
}

func (c *CalendarStorage) Range(f func(key int, event models.EventT) bool) {
	c.m.RLock()
	defer c.m.RUnlock()
	for k, e := range c.data {
		if !f(k, e) {
			break
		}
	}
}
