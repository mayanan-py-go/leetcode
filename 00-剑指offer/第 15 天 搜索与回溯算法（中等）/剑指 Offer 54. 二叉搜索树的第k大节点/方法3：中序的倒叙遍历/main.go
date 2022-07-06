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

	fmt.Println(kthLargest(root, 1))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func kthLargest(root *TreeNode, k int) int {
	/*
	最坏时间复杂度：O(N)
	空间复杂度：O(N)
	 */
	var res int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		if k == 0 {
			return
		}
		k--
		if k == 0 {
			res = node.Val
		}
		dfs(node.Left)
	}

	dfs(root)
	return res
}













