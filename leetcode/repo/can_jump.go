package repo

func canJump(nums []int) bool {
	// 维护一个maxReach变量
	maxReach := 0
	for i := range nums {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, nums[i]+i)
		if maxReach > len(nums) {
			return true
		}
	}
	return true
}
