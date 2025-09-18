package repo

func maxCoins(nums []int) int {
	// dp[i][j]表示nums[i,j]内部的最大戳破硬币数量
	n := len(nums)
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	nums = append(append([]int{1}, nums...), 1)
	for length := 2; length < n+2; length++ {
		for i := 0; i+length < n+2; i++ {
			j := i + length
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+nums[k]*nums[i]*nums[j])
			}
		}
	}
	return dp[0][n+1]
}

/*

// 贪心是错误的，会误判
// 戳气球的收益跟“剩下哪些气球”有关，不是局部贪心就能解的。
// 排序戳小的气球，可能导致大的气球失去了高价值的相邻乘积机会。

// 双向链表 + 节点排序
type DListNode struct {
	Val  int
	Next *DListNode
	Prev *DListNode
}

func maxCoins(nums []int) int {
	// 贪心：从中间戳，从小的开戳
	// 回溯：2^n不适合
	var res int
	NodeArray := make([]*DListNode, 0)
	head := GetDoubleLinkList(nums)
	for cur := head; cur != nil; cur = cur.Next {
		NodeArray = append(NodeArray, cur)
	}
	first, last := NodeArray[0], NodeArray[len(NodeArray)-1]
	NodeArray = NodeArray[1:]
	NodeArray = NodeArray[:len(NodeArray)-1]
	sort.Slice(NodeArray, func(i, j int) bool {
		return NodeArray[i].Val < NodeArray[j].Val
	})
	for _, node := range NodeArray {
		res += node.Val * node.Prev.Val * node.Next.Val
		//移除队列
		node.Next.Prev = node.Prev
		node.Prev.Next = node.Next
	}
	res += (min(first.Val, last.Val) + 1) * max(first.Val, last.Val)
	return res
}

func GetDoubleLinkList(nums []int) *DListNode {
	dummy := &DListNode{}
	cur := dummy
	for _, n := range nums {
		node := &DListNode{Val: n}
		cur.Next = node
		node.Prev = cur
		cur = node
	}
	dummy.Next.Prev = nil
	head := dummy.Next
	return head
}

*/
