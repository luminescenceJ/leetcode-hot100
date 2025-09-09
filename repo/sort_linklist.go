package repo

// 归并排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	start, mid := divide(head)
	start = sortList(start)
	mid = sortList(mid)
	return merge(start, mid)
}

func divide(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, nil
	}
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if prev != nil {
		prev.Next = nil
	}
	return head, slow
}

func merge(h1, h2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for h1 != nil && h2 != nil {
		if h1.Val > h2.Val {
			cur.Next = h2
			h2 = h2.Next
		} else {
			cur.Next = h1
			h1 = h1.Next
		}
		cur = cur.Next
	}
	if h1 != nil {
		cur.Next = h1
	}
	if h2 != nil {
		cur.Next = h2
	}
	return dummy.Next
}
