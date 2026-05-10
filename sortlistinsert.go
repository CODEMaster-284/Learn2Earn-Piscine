package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	node := &NodeI{Data: data_ref}

	if l == nil || data_ref < l.Data {
		node.Next = l
		return node
	}

	current := l
	for current.Next != nil && current.Next.Data <= data_ref {
		current = current.Next
	}

	node.Next = current.Next
	current.Next = node

	return l
}
