package lru

import "sync"

type CurrentLRU struct {
	lru *Lru
	lock sync.Mutex
}

func (c *CurrentLRU)Put(k,v interface{}){
	c.lock.Lock()
	c.lru.Put(k,v)
	c.lock.Unlock()
}

func (c *CurrentLRU)Get(k interface{})(v interface{},ok bool){
	c.lock.Lock()
	v,ok = c.lru.Get(k)
	c.lock.Unlock()
	return
}
