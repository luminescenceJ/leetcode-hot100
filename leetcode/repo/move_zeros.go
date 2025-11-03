package repo

func moveZeroes(nums []int) {
	left := 0
	for i := 0; i < len(nums); i++ {
		nums[left] = nums[i]
		if nums[i] == 0 {
			continue
		}
		left++
	}
	for i := left; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroes2(nums []int) {
	done, cur, n := 0, 0, len(nums)
	for cur < n {
		if nums[cur] != 0 {
			nums[done], nums[cur] = nums[cur], nums[done]
			done++
		}
		cur++
	}
}
