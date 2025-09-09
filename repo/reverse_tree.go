package repo

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var left, right *TreeNode
	if root.Left != nil {
		left = invertTree(root.Left)
	}
	if root.Right != nil {
		right = invertTree(root.Right)
	}
	root.Left = right
	root.Right = left
	return root
}
