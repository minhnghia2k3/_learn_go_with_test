package main

import (
	"sync"
	"sync/atomic"
)

type Counter interface {
	Inc()
	Value() int64
}

type DefaultCounter struct {
	mu    sync.Mutex
	value int
}

func NewDefaultCounter() Counter {
	return &DefaultCounter{}
}

func (c *DefaultCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *DefaultCounter) Value() int64 {
	return int64(c.value)
}

type AtomicCounter struct {
	value atomic.Int64
}

func NewAtomicCounter() Counter {
	return &AtomicCounter{}
}

func (c *AtomicCounter) Inc() {
	c.value.Add(1)
}

func (c *AtomicCounter) Value() int64 {
	return c.value.Load()
}
