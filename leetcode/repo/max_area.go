package repo

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	var res int
	for left < right {
		res = max(res, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
