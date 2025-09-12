package repo

func pathSum(root *TreeNode, targetSum int) int {
	// 前缀优化
	// map[int]int表示这个节点之前是否的前缀和
	// curSum -prevSum  == targetSum => res++
	// O(N)
	hash := map[int]int{}
	hash[0] = 1
	var res int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, curSum int) {
		if node == nil {
			return
		}
		res += hash[curSum+node.Val-targetSum]
		hash[node.Val+curSum]++
		dfs(node.Left, node.Val+curSum)
		dfs(node.Right, node.Val+curSum)
		hash[node.Val+curSum]--
	}
	dfs(root, 0)
	return res
}

func pathSum2(root *TreeNode, targetSum int) int {
	// O(N^2)
	if root == nil {
		return 0
	}
	return rootSum(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func rootSum(root *TreeNode, target int) int {
	if root == nil {
		return 0
	}
	var res int
	if root.Val == target {
		res += 1
	}
	res += rootSum(root.Left, target-root.Val)
	res += rootSum(root.Right, target-root.Val)
	return res
}
