package repo

import "fmt"

func longestPalindrome(s string) string {
	res := string(s[0])
	var temp string
	for i := range s {
		if i != 0 {
			temp = fromCenter(s, i-1, i)
			if len(temp) > len(res) {
				res = temp
			}
		}
		temp = fromCenter(s, i, i)
		if len(temp) > len(res) {
			res = temp
		}
	}
	return res
}

func fromCenter(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
func TestLongestPalindrome() {
	str := "babad"
	fmt.Println(longestPalindrome(str))
}
