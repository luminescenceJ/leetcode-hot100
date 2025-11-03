package repo

func subsets(nums []int) [][]int {
	// 枚举index
	var res [][]int
	var dfs func(idx int)
	cur := []int{}
	dfs = func(idx int) {
		if idx == len(nums) {
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}
		cur = append(cur, nums[idx])
		dfs(idx + 1)
		cur = cur[:len(cur)-1]
		dfs(idx + 1)
	}
	dfs(0)
	return res
}
