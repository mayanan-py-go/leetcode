package main

import "fmt"

/*
给定一棵二叉搜索树，请找出其中第 k 大的节点的值。
https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/

示例 1:
 */


func main() {
	root := &TreeNode{Val: 3}
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 4}
	n3 := &TreeNode{Val: 2}

	root.Left, root.Right = n1, n2
	n1.Right = n3

	fmt.Println(kthLargest(root, 2))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func kthLargest(root *TreeNode, k int) int {
	/*
	时间复杂度：O(N)
	空间复杂度：O(N)
	 */
	var s1 = make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		s1 = append(s1, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return s1[len(s1)-k]
}













