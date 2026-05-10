package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}

	node := l
	for node != nil && pos > 0 {
		node = node.Next
		pos--
	}

	return node
}
