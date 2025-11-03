package sort

import "fmt"

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0 && nums[j+1] < nums[j]; j-- {
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
	}
}

func TestInsertSort() {
	nums := []int{5, 3, 4, 1, 2}
	insertSort(nums)
	fmt.Println(nums)
}
