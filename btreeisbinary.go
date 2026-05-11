package piscine

func BTreeIsBinary(root *TreeNode) bool {
	var prev string
	hasPrev := false

	var walk func(*TreeNode) bool
	walk = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if !walk(node.Left) {
			return false
		}

		if hasPrev && node.Data <= prev {
			return false
		}

		prev = node.Data
		hasPrev = true

		return walk(node.Right)
	}

	return walk(root)
}
