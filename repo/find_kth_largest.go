package repo

import "container/heap"

// 小根堆原理：元素在末尾发生处理
// 插入：append到数组末，然后依次和他的父节点(index-1)/2比较大小并交换
// 删除：头末节点交换，然后从头节点开始和孩子节点中更小的节点交换

// method1: implement Minority Heap
func findKthLargest1(nums []int, k int) int {
	// 优先级队列
	m := MinPile{
		item: make([]int, 0),
	}
	for i := range nums {
		if len(m.item) < k {
			m.Insert(nums[i])
		} else {
			if m.item[0] < nums[i] {
				m.Delete()
				m.Insert(nums[i])
			}
		}
	}
	return m.item[0]
}

// method2: with container/heap pkg
func findKthLargest2(nums []int, k int) int {

	h := make(Heap, 0)
	heap.Init(&h)
	for i := range nums {
		if h.Len() < k {
			heap.Push(&h, nums[i])
		} else {
			if nums[i] > h[0] {
				heap.Pop(&h)
				heap.Push(&h, nums[i])
			}
		}
	}
	return heap.Pop(&h).(int)
}

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] < h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
