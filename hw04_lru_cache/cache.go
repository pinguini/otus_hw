package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.Mutex
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	exists := false
	val, ok := c.items[key]
	if ok {
		c.queue.Remove(val)
		exists = true
	}

	c.items[key] = c.queue.PushFront(value)
	if c.queue.Len() > c.capacity {
		last := c.queue.Back()
		c.queue.Remove(last)
		for k := range c.items {
			if c.items[k] == last {
				delete(c.items, k)
			}
		}
	}
	c.mu.Unlock()
	return exists
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	val, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(val)
		c.mu.Unlock()
		return val.Value, true
	}
	c.mu.Unlock()
	return nil, false
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	for item := range c.items {
		c.queue.Remove(c.items[item])
		delete(c.items, item)
	}
	c.mu.Unlock()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
