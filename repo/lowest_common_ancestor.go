package repo

// leetcode 236.最近公共最先
// time: O(N)
// space: O(N)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p == root || q == root {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}
	if l == nil {
		return r
	}
	return l
}
