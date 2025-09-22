package repo

func subarraySum2(nums []int, k int) int {
	// 前缀和 + hash表
	hashMap := make(map[int]int)
	preSum := make([]int, len(nums)+1)
	hashMap[0] = 1

	// preSum[i]表示截止到nums[i]的和
	var res int
	for i, n := range nums {
		preSum[i+1] = n + preSum[i]
		if v, ok := hashMap[preSum[i+1]-k]; ok {
			res += v
		}
		hashMap[preSum[i+1]]++
	}
	return res
}

func subarraySum(nums []int, k int) int {
	// 前缀和
	preSum := make([]int, len(nums)+1)
	for i, n := range nums {
		preSum[i+1] = n + preSum[i]
	}
	// preSum[i]表示截止到nums[i]的和
	var res int
	for i := 1; i <= len(nums); i++ {
		for j := 0; j < i; j++ {
			if preSum[i]-preSum[j] == k {
				res++
			}
		}
	}
	return res
}
