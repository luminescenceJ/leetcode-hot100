package repo

func rotate(matrix [][]int) {
	// 上下翻转
	upAndDown(matrix)
	// 主对角线翻转
	mainDig(matrix)
}

func upAndDown(matrix [][]int) {
	n := len(matrix) - 1
	for i := 0; i <= n/2; i++ {
		matrix[i], matrix[n-i] = matrix[n-i], matrix[i]
	}
}

func mainDig(matrix [][]int) {
	for i := range matrix {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

/*
1 2 3
4 5 6
7 8 9
*/

/*
7 8 9
4 5 6
1 2 3
*/

/*
7 4 1
8 5 2
9 6 3
*/
