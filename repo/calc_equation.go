package repo

type UnionSet struct {
	parent map[string]string
	weight map[string]float64 // a / parent[a] = weight[a]
}

func (u *UnionSet) Find(a string) (string, float64) {
	if _, ok := u.parent[a]; !ok {
		return "", -1.0 // 变量不存在
	}
	if u.parent[a] == a {
		return a, 1.0
	}
	p, w := u.Find(u.parent[a])
	u.parent[a] = p
	u.weight[a] *= w
	return p, u.weight[a]
}

func (u *UnionSet) Union(a, b string, value float64) {
	// value = a / b
	rootA, weightA := u.Find(a)
	rootB, weightB := u.Find(b)
	if rootA == rootB {
		return
	}
	u.parent[rootA] = rootB
	u.weight[rootA] = weightB * value / weightA
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	u := UnionSet{
		parent: make(map[string]string),
		weight: make(map[string]float64),
	}

	// 初始化并合并集合
	for i, eq := range equations {
		a, b := eq[0], eq[1]
		if _, ok := u.parent[a]; !ok {
			u.parent[a] = a
			u.weight[a] = 1.0
		}
		if _, ok := u.parent[b]; !ok {
			u.parent[b] = b
			u.weight[b] = 1.0
		}
		u.Union(a, b, values[i])
	}

	// 处理查询
	res := make([]float64, len(queries))
	for i, q := range queries {
		a, b := q[0], q[1]
		rootA, weightA := u.Find(a)
		rootB, weightB := u.Find(b)
		if rootA == "" || rootB == "" || rootA != rootB {
			res[i] = -1.0
		} else {
			res[i] = weightA / weightB
		}
	}
	return res
}
