package repo

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	// 数字，括号
	var tempNum int
	var tempStr string
	var res string
	var stack []string
	for i := 0; i < len(s); {
		if s[i] <= '9' && s[i] >= '0' {
			// 数字
			for s[i] <= '9' && s[i] >= '0' {
				tempNum = tempNum*10 + int(s[i]-'0')
				i++
			}
			stack = append(stack, strconv.Itoa(tempNum))
			tempNum = 0
		} else if s[i] == '[' {
			stack = append(stack, string(s[i]))
			i++
		} else if s[i] == ']' {
			for len(stack) > 0 && stack[len(stack)-1] != "[" {
				tempStr = stack[len(stack)-1] + tempStr
				stack = stack[:len(stack)-1]
			}
			// 出栈'['
			stack = stack[:len(stack)-1]
			tempNum, _ = strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			stack = append(stack, strings.Repeat(tempStr, tempNum))
			tempNum = 0
			tempStr = ""
			i++
		} else {
			for i < len(s) && ((s[i] <= 'z' && s[i] >= 'a') || (s[i] <= 'Z' && s[i] >= 'A')) {
				tempStr += string(s[i])
				i++
			}
			stack = append(stack, tempStr)
			tempStr = ""
		}
		fmt.Println(stack)
	}
	for _, s := range stack {
		res += s
	}
	return res
}
