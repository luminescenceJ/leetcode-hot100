package repo

import "fmt"

func removeInvalidParentheses(s string) []string {
	count := CountLeastDelete(s)
	var res []string
	var dfs func(int, int, int)
	curByte := []byte{}
	dfs = func(index int, left int, curCount int) {
		if index == len(s) && left == 0 && curCount == 0 {
			res = append(res, string(curByte))
			return
		}
		if left < 0 || index == len(s) || curCount < 0 {
			return
		}
		if s[index] == '(' {
			curByte = append(curByte, s[index])
			dfs(index+1, left+1, curCount)
			curByte = curByte[:len(curByte)-1]
			dfs(index+1, left, curCount-1)
		} else if s[index] == ')' {
			if left > 0 {
				curByte = append(curByte, s[index])
				dfs(index+1, left-1, curCount)
				curByte = curByte[:len(curByte)-1]
			}
			dfs(index+1, left, curCount-1)
		} else {
			curByte = append(curByte, s[index])
			dfs(index+1, left, curCount)
			curByte = curByte[:len(curByte)-1]
		}
	}
	dfs(0, 0, count)
	// res去重
	hashMap := map[string]bool{}
	for _, s := range res {
		hashMap[s] = true
	}
	Final := []string{}
	for k := range hashMap {
		Final = append(Final, k)
	}
	return Final
}
func CountLeastDelete(s string) int {
	leastDelete := 0
	leftBra := 0
	for _, ch := range s {
		if ch == ')' {
			if leftBra > 0 {
				leftBra--
			} else {
				leastDelete++
			}
		} else if ch == '(' {
			leftBra++
		}
	}
	leastDelete += leftBra
	return leastDelete
}
func TestRemoveInvalidParentheses() {
	s := "()())()"
	fmt.Println(removeInvalidParentheses(s))
}
