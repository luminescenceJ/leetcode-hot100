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

type MinPile struct {
	item []int
}

func (m *MinPile) Insert(x int) {
	m.item = append(m.item, x)
	m.HeapifyUp(len(m.item) - 1)
}

func (m *MinPile) HeapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if m.item[index] < m.item[parent] { // 小根堆：子节点 < 父节点
			m.item[index], m.item[parent] = m.item[parent], m.item[index]
			index = parent
		} else {
			break
		}
	}
}

func (m *MinPile) Delete() int {
	if len(m.item) == 0 {
		panic("heap is empty")
	}
	root := m.item[0]
	last := len(m.item) - 1
	m.item[0] = m.item[last]
	m.item = m.item[:last]
	m.HeapifyDown(0)
	return root
}

func (m *MinPile) HeapifyDown(index int) {
	n := len(m.item)
	for {
		left, right := 2*index+1, 2*index+2
		smallest := index

		if left < n && m.item[left] < m.item[smallest] {
			smallest = left
		}
		if right < n && m.item[right] < m.item[smallest] {
			smallest = right
		}
		if smallest == index {
			break
		}
		m.item[index], m.item[smallest] = m.item[smallest], m.item[index]
		index = smallest
	}
}
