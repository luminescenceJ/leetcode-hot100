package repo

import "fmt"

/*
奇数个节点的链表，当你使用快慢指针找到中间节点时，快指针在最后一轮会跳过中间节点，而慢指针会停在这个中间节点上。
反转链表的后半部分，中间节点不会参与反转，只是“跳过”了。
假设链表是 [1, 2, 3, 2, 1]（奇数个节点）

slow 指向 1 → 2 → 3，最终停在中间的 3。
fast 指向 1 → 3 → 1，最终到达链表的末尾。

此时，slow 正好指向中间节点 3，此时链表被分成了两部分：
左半部分：[1, 2, 3]
右半部分：[2, 1],反转后得到 [1, 2]
对比时中间的奇数节点被忽略了
*/

func isPalindrome(head *ListNode) bool {
	newhead := findMid(head)
	newhead = reverse(newhead)
	for l1, l2 := head, newhead; l1 != nil && l2 != nil; {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}

func findMid(head *ListNode) *ListNode {
	// slow 作为第二条链表的起始节点
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if pre != nil {
		pre.Next = nil
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	var prev *ListNode
	for cur != nil {
		nxt := cur.Next
		cur.Next = prev
		prev = cur
		cur = nxt
	}
	return prev
}

func TestisPalindrome() {
	data := []int{1, 2, 2, 1}
	l := GenerateListNode(data)
	ans := isPalindrome(l)
	fmt.Println(ans)
}
