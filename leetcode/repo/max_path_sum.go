package repo

import "math"

func maxPathSum(root *TreeNode) int {
	var res int = math.MinInt32
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(0, dfs(root.Left))
		right := max(0, dfs(root.Right))
		res = max(res, root.Val+left+right)
		return root.Val + max(left, right)
	}
	dfs(root)
	return res
}
