package repo

import "math"

func maxProfit(prices []int) int {
	curLow := math.MaxInt32
	res := 0
	for _, n := range prices {
		if n < curLow {
			curLow = n
		}
		res = max(res, n-curLow)
	}
	return res
}
