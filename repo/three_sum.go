package repo

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	// 枚举起点，+1为左端点，某尾为右端点
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right]+nums[i] > 0 {
				right--
			} else if nums[left]+nums[right]+nums[i] < 0 {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for right > left && nums[right+1] == nums[right] {
					right--
				}
			}

		}
	}
	return res
}

func TestThreeSum() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)

}
