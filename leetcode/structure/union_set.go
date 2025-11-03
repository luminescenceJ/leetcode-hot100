package structure

import "math"

type UnionSet struct {
	parent map[int]int
}

func NewUnionSet(arr []int) *UnionSet {
	us := &UnionSet{make(map[int]int)}
	for i := 0; i < len(arr); i++ {
		us.parent[arr[i]] = i
	}
	return us
}
func (us *UnionSet) Find(x int) int {
	if _, ok := us.parent[x]; !ok {
		return math.MinInt32
	}
	if x != us.parent[x] {
		x = us.Find(us.parent[x])
	}
	return us.parent[x]
}

func (us *UnionSet) Union(i, j int) bool {
	pi := us.Find(i)
	pj := us.Find(j)
	if pi == pj {
		return true
	}
	us.parent[pi] = pj
	return true
}

func (us *UnionSet) IsSameSet(i, j int) bool {
	if _, ok := us.parent[i]; !ok {
		return false
	}
	if _, ok := us.parent[j]; !ok {
		return false
	}
	return us.Find(i) == us.Find(j)
}
