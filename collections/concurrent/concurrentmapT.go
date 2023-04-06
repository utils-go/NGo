package concurrent

import "sync"

type ConcurrentMap[TKey comparable, TValue any] struct {
	mu  sync.RWMutex
	dic map[TKey]TValue
}

func NewCurrentMap[Tkey comparable, TValue any]() *ConcurrentMap[Tkey, TValue] {
	return &ConcurrentMap[Tkey, TValue]{
		dic: make(map[Tkey]TValue),
	}
}

func (c *ConcurrentMap[TKey, TValue]) Get(key TKey) (TValue, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if v, ok := c.dic[key]; ok {
		return v, true
	} else {
		return v, false
	}
}
func (c *ConcurrentMap[TKey, TValue]) Set(key TKey, value TValue) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.dic[key] = value
}
