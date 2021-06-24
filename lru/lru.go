package lru

import (
	"fmt"
	"github.com/seeadoog/goutils/excp"
)

type SizeOf func(v interface{}) int

type Lru struct {
	list   *List
	data   map[interface{}]*Node
	cap    int
	size   int
	sizeOf SizeOf
}

func NewLru(cap int,sizeof SizeOf)*Lru{
	return &Lru{
		list:   newList(),
		data: map[interface{}]*Node{},
		cap:    cap,
		size:   0,
		sizeOf: sizeof,
	}
}

func (l *Lru) moveToBack(n *Node) {
	l.list.Remove(n)
	l.list.PushBack(n)
}

func (l *Lru) Put(k, v interface{}) {
	size := l.sizeOf(v)
	if size > l.cap{
		excp.Throw(fmt.Errorf("size of %d > cap of lru: %d",size,l.cap))
	}

	node, ok := l.data[k]
	if ok {
		node.Val = v
		l.list.Remove(node)
		l.size -= l.sizeOf(v)
	}else{
		node = &Node{
			Key:  k,
			Val:  v,
			List: l.list,
		}
		l.data[k] = node
	}
	// cap is not full
	l.list.PushBack(node)
	l.size += size

	for l.size > l.cap{
		n := l.list.Front()
		if n == nil{
			excp.Throw(fmt.Errorf("front at nil node of lru list"))
		}
		l.list.Remove(n)
		delete(l.data, n.Key)
		l.size -= l.sizeOf(n.Val)
	}

}

func (l *Lru)Get(k interface{})(interface{},bool){
	node,ok := l.data[k]
	if !ok{
		return nil,false
	}
	l.moveToBack(node)
	return node.Val,true
}

func (l *Lru)Size()int{
	return l.size
}

func (l *Lru)Cap()int{
	return l.cap
}
