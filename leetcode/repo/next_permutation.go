package repo

import "sort"

func nextPermutation(nums []int) {
	// 倒数找到第一个顺序对，然后交换顺序
	// 如果不存在顺序对直接排序
	var i int
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			// 出现顺序对,找到大于当前位置的最小值，然后排序剩余元素
			curMin := i + 1
			for j := i + 1; j < len(nums); j++ {
				if nums[j] > nums[i] && nums[curMin] > nums[j] {
					curMin = j
				}
			}
			nums[i], nums[curMin] = nums[curMin], nums[i]
			break
		}
	}
	if i < 0 {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		return
	}
	sort.Ints(nums[i+1:])
}
