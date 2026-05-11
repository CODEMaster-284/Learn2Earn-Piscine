package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	min := root

	leftMin := BTreeMin(root.Left)
	if leftMin != nil && leftMin.Data < min.Data {
		min = leftMin
	}

	rightMin := BTreeMin(root.Right)
	if rightMin != nil && rightMin.Data < min.Data {
		min = rightMin
	}

	return min
}
