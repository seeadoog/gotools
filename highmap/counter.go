package highmap

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	hm *HighMap
	tags []string
	mu sync.Mutex
}
func NewCounter(tags ...string)*Counter{
	return &Counter{hm: NewHighMap(),tags:tags }
}

func (c *Counter)IncKeys(keys ...Tag){
	key := Keys(keys)
	var p *int64
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, val := range c.hm.Get(keys...) {
		if val.Key == key{
			p = val.V.(*int64)
			break
		}
	}
	if p == nil{
		p = new(int64)
		c.hm.Set(p,keys...)
	}
	atomic.AddInt64(p,1)
}

func (c *Counter)Inc(vals ...string){
	ks := make([]Tag, len(c.tags))
	for i, tag := range c.tags {
		ks[i] = Tag{Key: tag, Val: vals[i]}
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	cter  := c.hm.Get(ks...)

	if len(cter) == 0{
		counter := new(int64)
		*counter = 1

		c.hm.Set(counter,ks...)
	}else{
		atomic.AddInt64(cter[0].V.(*int64),1)
	}
}

func (c *Counter)GetCount(ks ...Tag)int{
	c.mu.Lock()
	defer c.mu.Unlock()
	sum := 0
	for _, v := range c.hm.Get(ks...) {
		sum +=int(*(v.V.(*int64)))
	}
	return sum
}

func (c *Counter)Show(){
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, val := range c.hm.indexes {
		for k, m := range val {
			fmt.Println(key,k,m)
		}
	}
}

