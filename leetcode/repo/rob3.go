package repo

func rob3(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		leftRob, leftNotRob := dfs(node.Left)
		rightRob, rightNotRob := dfs(node.Right)

		// 偷当前节点
		robCur := node.Val + leftNotRob + rightNotRob
		// 不偷当前节点
		notRobCur := max(leftRob, leftNotRob) + max(rightRob, rightNotRob)

		return robCur, notRobCur
	}

	robRoot, notRobRoot := dfs(root)
	return max(robRoot, notRobRoot)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
