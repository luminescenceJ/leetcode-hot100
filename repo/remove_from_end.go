package repo

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := head, dummy
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	if n == 1 && head.Next == nil {
		return nil
	}
	count := n
	for cur := head; cur != nil; cur = cur.Next {
		count--
	}
	if count == 0 {
		return head.Next
	}
	fast, slow := head, head
	for fast != nil && n > 0 {
		fast = fast.Next
		n--
	}
	var prev *ListNode
	for fast != nil {
		prev = slow
		fast = fast.Next
		slow = slow.Next
	}
	// slow为需要删除的节点
	if prev != nil && prev.Next != nil {
		prev.Next = prev.Next.Next
	}
	return head

}
