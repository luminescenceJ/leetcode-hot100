package repo

type Trie struct {
	node  [26]*Trie
	isEnd bool
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, w := range word {
		idx := int(w - 'a')
		if cur.node[idx] == nil {
			cur.node[idx] = &Trie{}
		}
		cur = cur.node[idx]
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, w := range word {
		if cur.node[int(w-'a')] == nil {
			return false
		}
		cur = cur.node[int(w-'a')]
	}
	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, w := range prefix {
		if cur.node[int(w-'a')] == nil {
			return false
		}
		cur = cur.node[int(w-'a')]
	}
	return true
}
