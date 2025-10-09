package repo

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 中位数：数组中第len(arr)/2-1的位置的数
	l1, l2 := len(nums1), len(nums2)
	size := l1 + l2
	left, right := 0, 0
	if size%2 == 0 {
		left, right = size/2-1, size/2
		return (float64(getKth(nums1, nums2, left+1)) + float64(getKth(nums1, nums2, right+1))) / 2
	} else {
		left, right = size/2, size/2
		return float64(getKth(nums1, nums2, left+1))
	}
	// 找到数组中的left,right的位置
	// 合并数组 + 排序 时间复杂度O(NlogN)
	// nums := append(nums1,nums2...)
	// sort.Ints(nums)
	// return (float64(nums[left]) + float64(nums[right]))/2
}

func getKth(nums1, nums2 []int, k int) int {
	// index1表示在nums1中的位置
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}

		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIdx1 := min(index1+half, len(nums1)) - 1
		newIdx2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIdx1], nums2[newIdx2]
		if pivot1 <= pivot2 {
			// pivot左侧的全部数排除
			k = k - (newIdx1 - index1 + 1)
			index1 = newIdx1 + 1
		} else {
			k -= newIdx2 - index2 + 1
			index2 = newIdx2 + 1
		}

	}
}
