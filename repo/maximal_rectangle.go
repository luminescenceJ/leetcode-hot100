package repo

func maximalRectangle(matrix [][]byte) int {
	var res int
	curHeight := make([]int, len(matrix[0]))

	for i := len(matrix) - 1; i >= 0; i-- {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				curHeight[j] += 1
			} else {
				curHeight[j] = 0
			}
		}
		res = max(largestRectangleArea(curHeight), res)
	}
	return res
}
