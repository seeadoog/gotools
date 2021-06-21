package highmap

import "sync"

type Gauge struct {
	hm *HighMap
	mu sync.RWMutex
}

func NewGauge()*Gauge{
	return &Gauge{hm: NewHighMap()}
}

func (c *Gauge)SetKvs(v float64,kvs ...string){
	c.Set(v, NKeys(kvs...)...)
}

func (c *Gauge) Set(v float64, tags ...Tag) {
	key := Keys(tags)
	c.mu.RLock()
	for _, val := range c.hm.Get(tags...) {
		if val.Key == key {
			val.V = v
			c.mu.RUnlock()
			return
		}
	}
	c.mu.RUnlock()

	c.mu.Lock()
	c.hm.Set(v, tags...)
	c.mu.Unlock()
}

func (c *Gauge) Count(tags ...Tag) (sum float64) {
	c.mu.RLock()
	for _, val := range c.hm.Get(tags...) {
		sum += val.V.(float64)
	}
	c.mu.RUnlock()
	return
}
