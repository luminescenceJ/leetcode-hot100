package repo

func combinationSum(candidates []int, target int) [][]int {
	var dfs func(index int, target int)
	res := [][]int{}
	cur := []int{}
	dfs = func(index int, target int) {
		if index == len(candidates) && target == 0 {
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}
		if index == len(candidates) || target < 0 {
			return
		}

		cur = append(cur, candidates[index])
		dfs(index, target-candidates[index])
		cur = cur[:len(cur)-1]

		dfs(index+1, target)
	}
	dfs(0, target)
	return res
}
