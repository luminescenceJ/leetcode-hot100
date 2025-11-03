package repo

func singleNumber(nums []int) int {
	var res int
	for _, n := range nums {
		res ^= n
	}
	return res
}
