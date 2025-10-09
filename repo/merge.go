package repo

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	stack := []int{}
	for i := range intervals {
		if len(stack) == 0 {
			stack = append(stack, intervals[i]...)
			continue
		}
		if stack[1] >= intervals[i][0] {
			// 可以连接
			stack = []int{min(stack[0], intervals[i][0]), max(intervals[i][1], stack[1])}
		} else {
			// 不能连接
			res = append(res, stack)
			stack = intervals[i]
		}
	}
	if len(stack) > 0 {
		res = append(res, stack)
	}
	return res
}
