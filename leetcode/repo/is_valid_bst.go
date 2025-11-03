package repo

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isValidBST(root.Left) || !isValidBST(root.Right) {
		return false
	}
	if root.Left != nil {
		_, leftMax := dfs_(root.Left)
		if root.Val <= leftMax {
			return false
		}
	}
	if root.Right != nil {
		rightMin, _ := dfs_(root.Right)
		if root.Val >= rightMin {
			return false
		}
	}
	return true
}

func dfs_(root *TreeNode) (int, int) {
	if root == nil {
		return math.MaxInt32, math.MinInt32
	}
	leftMin, leftMax := dfs_(root.Left)
	rightMin, rightMax := dfs_(root.Right)
	return min(min(root.Val, leftMin), rightMin), max(max(root.Val, leftMax), rightMax)
}
