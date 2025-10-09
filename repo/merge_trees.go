package repo

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 层序遍历
	// 递归：后序遍历
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	node := &TreeNode{}
	node.Left = mergeTrees(root1.Left, root2.Left)
	node.Right = mergeTrees(root1.Right, root2.Right)
	node.Val = root1.Val + root2.Val
	return node
}
