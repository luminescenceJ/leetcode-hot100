package repo

func canFinish(numCourses int, prerequisites [][]int) bool {
	prevMap := map[int][]int{}
	// 寻找是否存在环,三色判断, 未学0，正在学1，已学2
	for _, item := range prerequisites {
		prevMap[item[0]] = append(prevMap[item[0]], item[1])
	}
	array := make([]int, numCourses)
	var dfs func(int) bool
	dfs = func(index int) bool {
		if array[index] == 1 {
			return false
		}
		if array[index] == 2 {
			return true
		}
		array[index] = 1
		for _, k := range prevMap[index] {
			if dfs(k) != true {
				return false
			}
		}
		array[index] = 2
		return true
	}
	for _, item := range prerequisites {
		if !dfs(item[0]) {
			return false
		}
	}
	return true
}
