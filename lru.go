package cache

import (
	"container/list"
	"sync"
)

type CacheElement interface {
	Key() string
}

type LRUCache struct {
	m        map[string]*list.Element
	l        *list.List
	capacity int
	mutex    sync.RWMutex
}

func NewLRUCache(capacity int) *LRUCache {
	if capacity <= 0 {
		return nil
	} else {
		return &LRUCache{
			m:        make(map[string]*list.Element, capacity),
			l:        list.New(),
			capacity: capacity,
		}
	}
}

func (c *LRUCache) See(el CacheElement) {
	if el == nil {
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if listEl, exists := c.m[el.Key()]; exists {
		listEl.Value = el
		c.l.MoveToFront(listEl)
	} else {
		if c.l.Len() < c.capacity {
			listEl := c.l.PushFront(el)
			c.m[el.Key()] = listEl
		} else {
			tail := c.l.Back()
			if tailVal, ok := tail.Value.(CacheElement); ok {
				tail.Value = el
				c.l.MoveToFront(tail)
				delete(c.m, tailVal.Key())
				c.m[el.Key()] = tail
			}
		}
	}
}

func (c *LRUCache) Get(key string) (el CacheElement) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if el, exists := c.m[key]; exists {
		if cacheEl, ok := el.Value.(CacheElement); ok {
			return cacheEl
		}
	}

	return nil
}