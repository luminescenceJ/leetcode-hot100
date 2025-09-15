package repo

func longestConsecutive(nums []int) int {
	// hashMap 存储nums<bool>
	hash := make(map[int]bool)
	for _, n := range nums {
		hash[n] = true
	}
	var res int
	for n, _ := range hash {
		if hash[n-1] {
			// not best answer
			continue
		}
		var tempRes int = 1
		var ok bool
		for _, ok = hash[n+1]; ok; {
			tempRes++
			n++
			_, ok = hash[n+1]
		}
		res = max(res, tempRes)
	}
	return res
}
