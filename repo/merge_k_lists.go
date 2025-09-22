package repo

import (
	"container/heap"
)

type queue struct {
	item []*ListNode
}

func (q *queue) Len() int {
	return len(q.item)
}
func (q *queue) Less(i, j int) bool {
	return q.item[i].Val < q.item[j].Val
}
func (q *queue) Swap(i, j int) {
	q.item[i], q.item[j] = q.item[j], q.item[i]
}
func (q *queue) Pop() interface{} {
	v := q.item[q.Len()-1]
	q.item = q.item[:q.Len()-1]
	return v
}
func (q *queue) Push(x interface{}) {
	q.item = append(q.item, x.(*ListNode))
}
func mergeKLists(lists []*ListNode) *ListNode {
	// 小根堆
	item := []*ListNode{}
	for i := range lists {
		if lists[i] == nil {
			continue
		}
		item = append(item, lists[i])
	}

	dummy := &ListNode{}
	cur := dummy

	q := queue{item: item}
	heap.Init(&q)
	for q.Len() > 0 {
		node := heap.Pop(&q).(*ListNode)
		if node == nil {
			continue
		}
		cur.Next = node
		cur = cur.Next
		if node.Next != nil {
			heap.Push(&q, node.Next)
		}
	}
	return dummy.Next
}
