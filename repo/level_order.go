package repo

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := []int{}
		var newQueue []*TreeNode
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				continue
			}
			cur = append(cur, node.Val)
			newQueue = append(newQueue, node.Left, node.Right)
		}
		if len(cur) > 0 {
			res = append(res, cur)
		}

		queue = newQueue
	}
	return res
}
