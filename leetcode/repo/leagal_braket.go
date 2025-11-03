package repo

func isValid(s string) bool {
	stack := []rune{}
	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, ch)
		}
		if ch == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		}

		if ch == '[' {
			stack = append(stack, ch)
		}
		if ch == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}

		if ch == '{' {
			stack = append(stack, ch)
		}
		if ch == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
