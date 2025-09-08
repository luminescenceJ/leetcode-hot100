package repo

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummy := &ListNode{}
	cur := dummy
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func PrintListNode(node *ListNode) {
	if node == nil {
		return
	}
	cur := node
	for cur != nil {
		fmt.Printf("%d,", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("\n")
}
