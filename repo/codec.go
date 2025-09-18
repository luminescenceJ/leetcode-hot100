package repo

// 层序遍历法

//import (
//	"fmt"
//	"strconv"
//	"strings"
//)
//
//type Codec struct {
//}
//
//// Serializes a tree to a single string.
//func (this *Codec) serialize(root *TreeNode) string {
//	if root == nil {
//		return ""
//	}
//	// 层序遍历
//	var res string
//	queue := []*TreeNode{root}
//
//	curHashNode := true
//	for len(queue) > 0 && curHashNode {
//		curHashNode = false
//		newQueue := []*TreeNode{}
//		for len(queue) > 0 {
//			node := queue[0]
//			queue = queue[1:]
//			if node == nil {
//				newQueue = append(newQueue, nil)
//				newQueue = append(newQueue, nil)
//				res += "null "
//				continue
//			}
//			res += strconv.Itoa(node.Val) + " "
//			if node.Left != nil || node.Right != nil {
//				curHashNode = true
//			}
//			newQueue = append(newQueue, node.Left)
//			newQueue = append(newQueue, node.Right)
//		}
//		queue = newQueue
//	}
//	return res
//}
//
//// Deserializes your encoded data to tree.
//func (this *Codec) deserialize(data string) *TreeNode {
//	array := strings.Split(strings.TrimSpace(data), " ")
//	dummy := &TreeNode{}
//	var dfs func(*TreeNode, int) *TreeNode
//	dfs = func(cur *TreeNode, index int) *TreeNode {
//		if index >= len(array) || array[index] == "null" {
//			return nil
//		}
//		val, _ := strconv.Atoi(array[index])
//		cur = &TreeNode{Val: val}
//		cur.Left = dfs(cur.Left, 2*index+1)
//		cur.Right = dfs(cur.Right, 2*index+2)
//		return cur
//	}
//
//	return dfs(dummy, 0)
//}
//
//func TestCodec() {
//	codec := Codec{}
//	str := "1 2 3 4 5 6 "
//	root := codec.deserialize(str)
//	fmt.Println(root)
//	newStr := codec.serialize(root)
//	fmt.Println(newStr)
//}

import (
	"strconv"
	"strings"
)

// DFS

type Codec struct {
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var dfs func(*TreeNode)
	var res = ""
	dfs = func(node *TreeNode) {
		if node == nil {
			res += "null "
			return
		}
		res += strconv.Itoa(node.Val) + " "
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	array := strings.Split(strings.Trim(data, " "), " ")
	var build func() *TreeNode
	idx := 0
	build = func() *TreeNode {
		if idx >= len(array) {
			return nil
		}
		if array[idx] == "null" {
			idx++
			return nil
		}
		val, _ := strconv.Atoi(array[idx])
		cur := &TreeNode{Val: val}
		idx++
		cur.Left = build()
		cur.Right = build()

		return cur
	}
	return build()
}
