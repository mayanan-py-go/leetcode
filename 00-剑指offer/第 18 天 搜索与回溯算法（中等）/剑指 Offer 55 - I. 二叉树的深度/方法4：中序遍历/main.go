package main

import (
	"fmt"
)

/*
输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。
https://leetcode.cn/problems/er-cha-shu-de-shen-du-lcof/
例如：

给定二叉树 [3,9,20,null,null,15,7]，
*/

func main() {
	root := &TreeNode{Val: 3}
	n1 := &TreeNode{Val: 9}
	n2 := &TreeNode{Val: 20}
	n3 := &TreeNode{Val: 15}
	n4 := &TreeNode{Val: 7}
	root.Left, root.Right = n1, n2
	n2.Left, n2.Right = n3, n4
	fmt.Println(maxDepth(root))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func maxDepth(root *TreeNode) int {
	var res int
	var depth int
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		depth++
		traverse(node.Left)
		res = max(res, depth)
		traverse(node.Right)
		depth--
	}
	traverse(root)
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}













