package repo

func moveZeroes(nums []int) {
	done, cur, n := 0, 0, len(nums)
	for cur < n {
		if nums[cur] != 0 {
			nums[done], nums[cur] = nums[cur], nums[done]
			done++
		}
		cur++
	}
}
