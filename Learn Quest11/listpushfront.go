package piscine

func ListPushFront(l *List, data interface{}) {
	if l == nil {
		return
	}

	node := &NodeL{
		Data: data,
		Next: l.Head,
	}

	l.Head = node
	if l.Tail == nil {
		l.Tail = node
	}
}
