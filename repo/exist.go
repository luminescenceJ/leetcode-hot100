package repo

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	startRow, startCol := []int{}, []int{}
	for i := range board {
		for j := range board[0] {
			if board[i][j] == word[0] {
				startRow = append(startRow, i)
				startCol = append(startCol, j)
			}
		}
	}
	var dfs func(i, j int, idx int)
	visited := make([]bool, m*n)
	direct := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var res bool
	dfs = func(i, j int, idx int) {
		if idx == len(word) || res {
			res = true
			return
		}
		if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) && !visited[i*n+j] && board[i][j] == word[idx] {
			visited[i*n+j] = true
			for _, d := range direct {
				dfs(i+d[0], j+d[1], idx+1)
			}
			visited[i*n+j] = false
		}
	}
	for i := range startCol {
		dfs(startRow[i], startCol[i], 0)
	}
	return res
}
