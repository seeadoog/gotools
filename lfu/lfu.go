package lfu

type LFU struct {
	fre map[int]*List
	data map[interface{}]*Node
	size int
	cap int
	minFrequent int
}

func NewLFU(cap int)*LFU{
	return &LFU{
		fre: map[int]*List{},
		data: map[interface{}]*Node{},
		size:        0,
		cap:         cap,
		minFrequent: 0,
	}
}

func (l *LFU)setFre(frequent int,node *Node){
	if l.fre[frequent] == nil{
		l.fre[frequent] = newList()
	}
	l.fre[frequent].PushBack(node)
}

func (l *LFU)removeFre(frequent int,node *Node){
	list ,ok := l.fre[frequent]
	if ok{
		list.Remove(node)
	}
}

func (l *LFU) Put(k,v interface{})  {
	if l.cap == 0{
		return
	}
	node,ok := l.data[k]
	if ok{
		l.removeFre(node.Frequent,node)
		node.Frequent ++
		node.Val = v
		l.setFre(node.Frequent,node)
		if l.minFrequent == node.Frequent-1{
			// 当最小的frequent 列表为空时，表明最小的使用频率发生了变化，变成当前的使用频率
			if l.fre[l.minFrequent].Size() == 0{
				l.minFrequent ++
			}
		}
		return
	}
	node = &Node{
		Key:      k,
		Val:      v,
		Frequent: 1,
	}
	l.setFre(node.Frequent,node)
	l.data[k] = node
	l.size ++

	if l.size <= l.cap{
		l.minFrequent = 1
		return
	}
	l.size --

	// 删除最近最少使用的
	list  := l.fre[l.minFrequent]
	n := list.Front()
	delete(l.data,n.Key)
	l.removeFre(n.Frequent,n)
	l.minFrequent = 1 // 有新的元素加入，最小frequent一定是1
	if list.Size() == 0{
		delete(l.fre,n.Frequent)
	}

}



func (l *LFU)Get(k interface{})(interface{},bool){
	node ,ok := l.data[k]
	if !ok{
		return nil,false
	}
	l.removeFre(node.Frequent,node)
	node.Frequent ++
	l.setFre(node.Frequent,node)
	// 获取的是最小的frequent
	if l.minFrequent == node.Frequent-1{
		// 当最小的frequent 列表为空时，表明最小的使用频率发生了变化，变成当前的使用频率
		if l.fre[l.minFrequent].Size() == 0{
			l.minFrequent ++
		}
	}
	return node.Val,true
}



type Node struct {
	Next *Node
	Pre *Node
	Key interface{}
	Val interface{}
	List *List
	Frequent int
}

func newList()*List{
	l := new(List)
	l.back = new(Node)
	l.front = new(Node)
	l.front.Next = l.back
	l.back.Pre = l.front
	l.front.List = l
	l.back.List = l
	return l
}


type List struct {
	front *Node
	back *Node
	size int
}

func (l *List)PushBack(n *Node){
	backPre := l.back.Pre
	backPre.Next = n
	n.Next = l.back
	l.back.Pre = n
	n.Pre = backPre
	l.size ++
}

func (l *List)Remove(n *Node){
	n.Pre.Next = n.Next
	n.Next.Pre = n.Pre
	l.size --
}

func (l *List)Front()*Node{
	if l.front.Next == l.back{
		return nil
	}
	return l.front.Next
}

func (l *List)Back()*Node{
	if l.back.Pre == l.front{
		return nil
	}
	return l.back.Pre
}

func (l *List)MoveToHead(n *Node){
	l.Remove(n)
	l.PushBack(n)
}

func (l *List)Size()int{
	return l.size
}
