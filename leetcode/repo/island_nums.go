package repo

// dfs扩散
func numIslands(grid [][]byte) int {
	visit := make([]bool, len(grid)*len(grid[0]))
	w := len(grid[0])
	var res int
	var dir = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i >= len(grid) || i < 0 || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
			return
		}
		if visit[i*w+j] {
			return
		}
		visit[i*w+j] = true
		for _, d := range dir {
			dfs(i+d[0], j+d[1])
		}
		return
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' && !visit[i*w+j] {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}
