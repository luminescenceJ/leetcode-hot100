package repo

func sortColors(nums []int) {
	left, right := -1, len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			left++
		} else if nums[i] == 2 {
			right--
		}
	}
	for i := 0; i <= left; i++ {
		nums[i] = 0
	}
	for i := left + 1; i < right; i++ {
		nums[i] = 1
	}
	for i := right; i < len(nums); i++ {
		nums[i] = 2
	}
}
