package repo

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	var dfs func(*TreeNode) int
	// dfs(root)计算以root为头的最长路径
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		res = max(res, left+right+2)
		return max(left, right) + 1
	}
	dfs(root)
	return res
}
