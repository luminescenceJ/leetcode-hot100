package repo

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 && matrix[i][j] != target {
		if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return i < len(matrix) && j >= 0 && matrix[i][j] == target
}
