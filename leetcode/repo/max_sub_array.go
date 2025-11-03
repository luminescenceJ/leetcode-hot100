package repo

import "math"

func maxSubArray(nums []int) int {
	// 枚举每个元素作为终点的最大和
	// curSum := max(nums[i],curSum + nums[i])
	res := math.MinInt32
	var curSum = 0
	for i := range nums {
		curSum = max(curSum, 0) + nums[i]
		res = max(res, curSum)
	}
	return res
}
