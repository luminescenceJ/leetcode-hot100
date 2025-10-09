package repo

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num, res := 0, 0
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil || l2 != nil || res != 0 {
		num = 0
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}
		num += res
		res = num / 10
		num = num % 10
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return dummy.Next
}
