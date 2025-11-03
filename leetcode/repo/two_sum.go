package repo

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, n := range nums {
		if idx, ok := hashMap[target-n]; ok {
			return []int{idx, i}
		}
		hashMap[n] = i
	}
	return []int{}
}
