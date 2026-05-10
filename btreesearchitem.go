package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	current := root

	for current != nil {
		if elem < current.Data {
			current = current.Left
		} else if elem > current.Data {
			current = current.Right
		} else {
			return current
		}
	}

	return nil
}
