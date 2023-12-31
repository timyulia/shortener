package cache

import (
	"sync"
)

type InMemory struct {
	pair  map[string]string
	mutex *sync.RWMutex
}

func New() *InMemory {
	return &InMemory{make(map[string]string), &sync.RWMutex{}}
}

func (c *InMemory) GetLongURL(short string) (string, error) {
	c.mutex.RLock()
	long := c.pair[short]
	c.mutex.RUnlock()
	return long, nil
}

func (c *InMemory) AddLinksPair(short, long string) error {
	c.mutex.Lock()
	c.pair[short] = long
	c.mutex.Unlock()
	return nil
}
