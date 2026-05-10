package piscine

func ListReverse(l *List) {
	if l == nil || l.Head == nil {
		return
	}

	oldHead := l.Head
	var prev *NodeL
	current := l.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev
	l.Tail = oldHead
}
