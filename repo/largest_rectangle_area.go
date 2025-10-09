package repo

func largestRectangleArea(heights []int) int {
	// 枚举每一个柱子的高度和左右边界
	fromLeft := make([]int, len(heights))
	stack := []int{}
	for i, h := range heights {
		// 保留单调递增的栈
		for len(stack) > 0 && h <= heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			fromLeft[i] = 0
		} else {
			fromLeft[i] = stack[len(stack)-1] + 1
		}
		stack = append(stack, i)
	}

	fromRight := make([]int, len(heights))
	stack = []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			fromRight[i] = len(heights) - 1
		} else {
			fromRight[i] = stack[len(stack)-1] - 1
		}
		stack = append(stack, i)
	}
	var res int

	for i, h := range heights {
		res = max(res, h*(fromRight[i]-fromLeft[i]+1))
	}
	return res
}
