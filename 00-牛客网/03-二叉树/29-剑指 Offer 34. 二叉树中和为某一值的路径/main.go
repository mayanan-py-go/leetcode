package main

import "fmt"

/*

 */

func main() {
	root := &TreeNode{Val: 1}
	n1 := &TreeNode{Val: 2}
	n2 := &TreeNode{Val: 3}
	n3 := &TreeNode{Val: 4}
	n4 := &TreeNode{Val: 5}
	root.Left, root.Right = n1, n2
	n1.Left, n1.Right = n3, n4
	fmt.Println(pathSum(root, 8))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func pathSum(root *TreeNode, target int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		dfs(node.Left)
		if node.Left == nil && node.Right == nil && sum(path) == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ret = append(ret, tmp)
		}
		dfs(node.Right)
		path = path[:len(path)-1]
	}
	dfs(root)
	return ret
}
func sum(li []int) (ret int) {
	for _, v := range li {
		ret += v
	}
	return ret
}









