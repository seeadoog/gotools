package lru

import (
	"fmt"
	"testing"
)

func TestLru_Get(t *testing.T) {
	l := NewLru(3, func(v interface{}) int {
		return v.(int)
	})

	l.Put(1,2)
	l.Put(2,2)
	l.Put(3,2)
	l.Put(4,1)

	fmt.Println(l.Get(4))
}
