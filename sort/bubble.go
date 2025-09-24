package sort

import "fmt"

// 冒泡排序:稳定排序
// 每一轮循环将最大的值不断交换到最后
func bubbleSort(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {

				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func TestBubbleSort() {
	nums := []int{3, 2, 1}
	bubbleSort1(nums)
	fmt.Println(nums)
}

// 剪枝优化
func bubbleSort1(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		swapped := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				swapped = true
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		if !swapped {
			break
		}
	}
}
