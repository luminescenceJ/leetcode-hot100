package repo

import (
	"fmt"
)

func reverseStr(s []byte) {
	// 将字符串整体翻转，再按照单词进行翻转
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	//fmt.Println(string(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			continue
		default:
			var j = i
			for j = i; j < len(s) && s[j] != ' '; j++ {
			}
			for i < j-1 {
				s[i], s[j-1] = s[j-1], s[i]
				i++
				j--
			}
		}
	}
	fmt.Println(string(s))
}

func TestReverseStr() {
	str := "the sky is blue"
	//reverseInPlace([]byte(str))
	reverseStr([]byte(str))
	//fmt.Println(str)
}

func reverseBytes(arr []byte, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func reverseInPlace(arr []byte) {
	reverseBytes(arr, 0, len(arr)-1)
	start := 0
	for i := 0; i <= len(arr); i++ {
		if i == len(arr) || arr[i] == ' ' {
			reverseBytes(arr, start, i-1)
			start = i + 1
		}
	}
	fmt.Println(string(arr))
}
