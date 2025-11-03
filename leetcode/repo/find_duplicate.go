package repo

func findDuplicate(nums []int) int {
	// 排序找 O(nlogn)

	// 占位法 O(n)

	// 双指针找环
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
