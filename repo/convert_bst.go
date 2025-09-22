package repo

func convertBST(root *TreeNode) *TreeNode {
	prefixSum := 0
	// 右中左
	var dfs func(*TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		node.Right = dfs(node.Right)
		node.Val += prefixSum
		prefixSum = node.Val
		node.Left = dfs(node.Left)
		return node
	}
	return dfs(root)
}
