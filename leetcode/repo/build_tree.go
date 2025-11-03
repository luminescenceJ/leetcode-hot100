package repo

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	mid := preorder[0]
	idx := -99999
	for i := range inorder {
		if inorder[i] == mid {
			idx = i
		}
	}
	// idx 为 中序遍历中根节点的下标
	inLeft := inorder[:idx]
	inRight := inorder[idx+1:]

	preLeft := preorder[1 : len(inLeft)+1]
	preRight := preorder[len(inLeft)+1:]

	node := &TreeNode{}
	node.Val = mid
	node.Left = buildTree(preLeft, inLeft)
	node.Right = buildTree(preRight, inRight)
	return node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// pre 中左右
	// in 左中右
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	mid := preorder[0]
	var index int
	for i, v := range inorder {
		if mid == v {
			index = i
			break
		}
	}
	// 中序遍历的根节点位置index
	root := &TreeNode{Val: mid}
	leftLen := index
	root.Left = buildTree(preorder[1:leftLen+1], inorder[:leftLen])
	root.Right = buildTree(preorder[leftLen+1:], inorder[leftLen+1:])
	return root
}
