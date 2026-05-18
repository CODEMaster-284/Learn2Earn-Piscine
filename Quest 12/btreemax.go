package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	max := root

	leftMax := BTreeMax(root.Left)
	if leftMax != nil && leftMax.Data > max.Data {
		max = leftMax
	}

	rightMax := BTreeMax(root.Right)
	if rightMax != nil && rightMax.Data > max.Data {
		max = rightMax
	}

	return max
}
