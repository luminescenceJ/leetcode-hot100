package repo

func trap(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	n := len(height)
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	var res int
	for i, h := range height {
		res += min(leftMax[i], rightMax[i]) - h
	}
	return res
}
