package repo

import "container/heap"

type item struct {
	val   int
	count int
}

type minHeap []item

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].count < h[j].count }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(item)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}
	for _, n := range nums {
		count[n]++
	}

	h := &minHeap{}
	heap.Init(h)
	for val, cnt := range count {
		if h.Len() < k {
			heap.Push(h, item{val, cnt})
		} else if cnt > (*h)[0].count {
			heap.Pop(h)
			heap.Push(h, item{val, cnt})
		}
	}

	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).(item).val
	}
	return res
}
