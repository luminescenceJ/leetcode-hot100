package repo

var arr = [][]byte{
	[]byte{},
	[]byte{},
	[]byte{'a', 'b', 'c'},
	[]byte{'d', 'e', 'f'},
	[]byte{'g', 'h', 'i'},
	[]byte{'j', 'k', 'l'},
	[]byte{'m', 'n', 'o'},
	[]byte{'p', 'q', 'r', 's'},
	[]byte{'t', 'u', 'v'},
	[]byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	var res []string
	var nums []int
	for _, ch := range digits {
		nums = append(nums, int(ch-'0'))
	}
	var dfs func(int)
	curCh := []byte{}
	dfs = func(index int) {
		if index == len(nums) {
			res = append(res, string(curCh))
			return
		}
		for _, ch := range arr[nums[index]] {
			curCh = append(curCh, ch)
			dfs(index + 1)
			curCh = curCh[:len(curCh)-1]
		}
	}
	dfs(0)
	return res
}
