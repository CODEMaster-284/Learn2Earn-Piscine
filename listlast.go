package piscine

func ListLast(l *List) interface{} {
	if l == nil || l.Head == nil {
		return nil
	}

	if l.Tail != nil {
		return l.Tail.Data
	}

	node := l.Head
	for node.Next != nil {
		node = node.Next
	}

	return node.Data
}
