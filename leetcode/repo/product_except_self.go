package repo

func productExceptSelf(nums []int) []int {
	// 前缀乘积和后缀乘积
	pre := make([]int, len(nums)+1)
	suf := make([]int, len(nums)+1)
	pre[0] = 1
	suf[len(nums)] = 1
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		pre[i+1] = pre[i] * nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i]
	}
	for i := range res {
		res[i] = pre[i] * suf[i+1]
	}
	return res
}
