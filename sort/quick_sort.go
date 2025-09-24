package sort

import "fmt"

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	l, r := left, right
	pivot := nums[l]
	for l < r {
		for l < r && nums[r] >= pivot {
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] <= pivot {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = pivot
	quickSort(nums, left, l-1)
	quickSort(nums, l+1, right)
}

func TestQuickSort() {
	nums := []int{9, 4, 6, 1, 3, 2, 8, 7, 6, 5}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
