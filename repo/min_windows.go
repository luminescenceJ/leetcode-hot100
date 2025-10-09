package repo

import (
	"fmt"
	"math"
)

func TestMinWindows() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}

func minWindow(s string, t string) string {
	// 滑动窗口
	tCounter := make(map[rune]int)
	for _, ch := range t {
		tCounter[ch]++
	}
	var resLen = math.MaxInt32
	var res = ""
	sCounter := make(map[rune]int)
	l, r := 0, 0
	for r < len(s) {
		sCounter[rune(s[r])]++
		r++
		for check2(tCounter, sCounter) {
			if resLen > r-l {
				resLen = r - l
				fmt.Println(resLen)
				res = s[l:r]
			}
			sCounter[rune(s[l])]--
			l++
		}
	}
	return res

}

func check2(tC map[rune]int, sC map[rune]int) bool {
	for i := range tC {
		if tC[i] > sC[i] {
			return false
		}
	}
	return true
}
