package lfu

import (
	"fmt"
	"testing"
)

func TestLFU_Put(t *testing.T) {
	lfu := NewLFU(2)
	lfu.Put(1,1)
	lfu.Put(2,2)
	lfu.Get(1)
	lfu.Get(1)
	lfu.Get(1)

	lfu.Put(3,1)

	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.Get(1))
}
