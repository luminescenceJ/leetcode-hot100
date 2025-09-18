package repo

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	// 单调栈 + 滑动窗口
	var monoStack []int
	for i := 0; i < k-1; i++ {
		for len(monoStack) > 0 && nums[i] > monoStack[len(monoStack)-1] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		monoStack = append(monoStack, nums[i])
	}
	var res []int
	for i := k - 1; i < len(nums); i++ {
		for len(monoStack) > 0 && nums[i] > monoStack[len(monoStack)-1] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		monoStack = append(monoStack, nums[i])
		res = append(res, monoStack[0])
		if nums[i-k+1] == monoStack[0] {
			monoStack = monoStack[1:]
		}
	}
	return res
}
func TestMaxSlidingWindow() {
	nums := []int{1, -1}
	k := 1
	fmt.Println(maxSlidingWindow(nums, k))
}
