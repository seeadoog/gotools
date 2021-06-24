package lru
/*

 */

type Node struct {
	Next *Node
	Pre *Node
	Key interface{}
	Val interface{}
	List *List
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
}

func (l *List)PushBack(n *Node){
	backPre := l.back.Pre
	backPre.Next = n
	n.Next = l.back
	l.back.Pre = n
	n.Pre = backPre
}

func (l *List)Remove(n *Node){
	n.Pre.Next = n.Next
	n.Next.Pre = n.Pre
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

