package highmap

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewGauge(t *testing.T) {
 	g:=NewGauge()
 	g.Set(5,NKey("name","xh"),NKey("age","15"))
 	g.Set(5,NKey("name","xh"),NKey("age","18"))
 	g.Set(5,NKey("name","xm"),NKey("age","18"))
 	g.Set(5,NKey("name","xh"))

 	fmt.Println(g.Count(NKey("name","xh")))
 	fmt.Println(g.Count(NKeys("age","18")...))
}

func BenchmarkLock(b *testing.B) {
	l := sync.RWMutex{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			l.RLock()
			l.RUnlock()
		}
	})
}
