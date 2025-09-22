package repo

func longestValidParentheses(s string) int {
	// 截止到i为止的最长字符串
	// stack + dp
	stack := []int{-1}
	// stack 中存储）的下标
	res := 0
	for i, ch := range s {
		if ch == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}

	return res
}
