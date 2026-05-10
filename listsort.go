package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	var sorted *NodeI

	for l != nil {
		next := l.Next

		if sorted == nil || l.Data < sorted.Data {
			l.Next = sorted
			sorted = l
		} else {
			current := sorted
			for current.Next != nil && current.Next.Data <= l.Data {
				current = current.Next
			}

			l.Next = current.Next
			current.Next = l
		}

		l = next
	}

	return sorted
}
