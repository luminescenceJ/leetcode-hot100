package repo

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}

	upper := upper_bound(nums, target)   // 第一个 > target
	lower := upper_bound(nums, target-1) // 第一个 > target-1, 即 >= target

	if lower < len(nums) && nums[lower] == target {
		res[0] = lower
	}
	if upper-1 >= 0 && nums[upper-1] == target {
		res[1] = upper - 1
	}
	return res
}

func upper_bound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
