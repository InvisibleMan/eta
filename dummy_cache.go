package main

import (
	"sync"
	"time"
)

type item struct {
	value      interface{}
	lastAccess int64
}

type DummyCache struct {
	m       map[string]*item
	l       sync.Mutex
	maxSize int
}

func NewDummyCache(size int, maxTTL int) (m *DummyCache) {
	m = &DummyCache{m: make(map[string]*item, size), maxSize: size}
	go func() {
		for now := range time.Tick(time.Second) {
			m.l.Lock()
			for k, v := range m.m {
				if now.Unix()-v.lastAccess > int64(maxTTL) {
					delete(m.m, k)
				}
			}
			m.l.Unlock()
		}
	}()
	return
}

func (m *DummyCache) Put(k string, v interface{}) {
	m.l.Lock()
	it, ok := m.m[k]
	if !ok {
		if len(m.m)+1 == m.maxSize {
			m.m = make(map[string]*item, m.maxSize)
		}

		it = &item{value: v}
		m.m[k] = it
	}
	it.lastAccess = time.Now().Unix()
	m.l.Unlock()
}

func (m *DummyCache) Get(k string) (interface{}, bool) {
	finded := false
	var v interface{}

	m.l.Lock()
	if it, ok := m.m[k]; ok {
		v = it.value
		finded = true
		it.lastAccess = time.Now().Unix()

	}
	m.l.Unlock()
	return v, finded
}
