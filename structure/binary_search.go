package structure

import "fmt"

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func searchLowerBound(nums []int, target int) int {
	// 找到大于等于target的最小下标位置
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func TestBinarySearch() {
	nums := []int{1, 3, 5, 6}
	target := 6
	fmt.Println(searchLowerBound(nums, target))
}
