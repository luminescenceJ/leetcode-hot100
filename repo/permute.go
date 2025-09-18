package repo

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			return
		}
		dfs(index + 1)
		// 将当前数字num插入到res
		if len(res) == 0 {
			res = [][]int{[]int{nums[index]}}
			return
		}
		new := [][]int{}
		for i := range res {
			// res[i]是待插入的数组
			// 遍历res[i]的每个可插入位置，插入得到新的res[i]

			temp := make([]int, len(res[i])+1)
			for j := 0; j <= len(res[i]); j++ {
				// 将nums[index]插入到第j个位置
				copy(temp[:j], res[i][:j])
				temp[j] = nums[index]
				copy(temp[j+1:], res[i][j:])
				new = append(new, temp)
				temp = make([]int, len(res[i])+1)
			}
		}
		res = new
	}
	dfs(0)
	return res
}

func permute2(nums []int) [][]int {
	var res [][]int
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
		}
		for i := index; i < len(nums); i++ {
			nums[i], nums[index] = nums[index], nums[i]
			dfs(index + 1)
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	dfs(0)
	return res
}

func TestPermute() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println(res)
}
