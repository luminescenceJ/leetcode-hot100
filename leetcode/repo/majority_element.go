package repo

func majorityElement(nums []int) int {
	// 波尔投票法
	voter := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == voter {
			count++
		} else {
			count--
		}
		if count == 0 {
			voter = nums[i]
			count = 1
		}
	}
	return voter
}
