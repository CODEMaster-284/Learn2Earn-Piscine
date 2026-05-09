package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	if l == nil {
		return
	}

	node := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	l.Tail.Next = node
	l.Tail = node
}
