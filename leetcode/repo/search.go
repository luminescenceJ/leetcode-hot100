package repo

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for mid := (right-left)/2 + left; left <= right; {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[right] {
			// 倒序在mid右侧,mid左侧递增
			if nums[mid] > target && nums[left] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// mid右侧递增
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		mid = (right-left)/2 + left
	}
	return -1
}
