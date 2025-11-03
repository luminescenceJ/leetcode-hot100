package repo

func findDisappearedNumbers(nums []int) []int {
	// 占位法
	for i := range nums {
		for nums[i] != i+1 {
			if nums[nums[i]-1] == nums[i] {
				break
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	var res []int
	for i, n := range nums {
		if n != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}
