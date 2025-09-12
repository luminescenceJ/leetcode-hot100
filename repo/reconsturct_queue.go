package repo

import "sort"

func reconstructQueue(people [][]int) [][]int {
	// 先按照身高排序，从高到低
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	// [7,0] [7,1] [6,1] [5,0] [5,2] [4,4]
	// 对于第i个元素，将他插入
	for i, peo := range people {
		if i == peo[1] {
			continue
		}
		temp := peo
		for j := i; j > peo[1]; j-- {
			people[j] = people[j-1]
		}
		people[peo[1]] = temp
	}
	return people

}
