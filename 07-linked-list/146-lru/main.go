package main

type LRU struct {
	capacity int
	size     int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

type Node struct {
	Val  int
	Key  int
	Next *Node
	Prev *Node
}

func NewLRU(capacity int) *LRU {
	l := &LRU{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
		head:     newNode(0, 0),
		tail:     newNode(0, 0),
	}
	l.head.Next = l.tail
	l.tail.Prev = l.head

	return l
}

func newNode(key, val int) *Node {
	return &Node{
		Val: val,
		Key: key,
	}
}

func (l *LRU) Get(key int) int {
	if v, ok := l.cache[key]; ok {
		l.moveToHead(v)
		return v.Val
	}

	return -1
}

func (l *LRU) Put(key, val int) {
	if v, ok := l.cache[key]; ok {
		v.Val = val
		l.moveToHead(v)
	} else {
		node := newNode(key, val)
		l.addToHead(node)
		l.size++
		if l.size > l.capacity {
			tail := l.deleteTail()
			delete(l.cache, tail.Key)
			l.size--
		}
	}
}

func (l *LRU) moveToHead(node *Node) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRU) addToHead(node *Node) {
	node.Next = l.head.Next
	node.Prev = l.head
	l.head.Next.Prev = node
	l.head.Next = node
}

func (l *LRU) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (l *LRU) deleteTail() *Node {
	tail := l.tail.Prev
	l.removeNode(tail)
	return tail
}

func main() {

}
