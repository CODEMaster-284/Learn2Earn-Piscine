package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	if node.Left == nil {
		return BTreeTransplant(root, node, node.Right)
	}

	if node.Right == nil {
		return BTreeTransplant(root, node, node.Left)
	}

	successor := node.Right
	for successor.Left != nil {
		successor = successor.Left
	}

	if successor.Parent != node {
		root = BTreeTransplant(root, successor, successor.Right)
		successor.Right = node.Right
		if successor.Right != nil {
			successor.Right.Parent = successor
		}
	}

	root = BTreeTransplant(root, node, successor)
	successor.Left = node.Left
	if successor.Left != nil {
		successor.Left.Parent = successor
	}

	return root
}
