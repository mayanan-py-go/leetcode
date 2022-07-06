package main

import (
	"fmt"
)

/*
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
叶子节点 是指没有子节点的节点。
https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
 */


func main() {
	root := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	root.Left, root.Right = n2, n3
	n2.Left, n2.Right = n4, n5

	fmt.Println(pathSum(root, 8))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func pathSum(root *TreeNode, target int) [][]int {
	var ans [][]int
	var path = make([]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() {path = path[:len(path)-1]}()  // 函数栈退出时，最后一个元素也同时删除
		if node.Left == nil && node.Right == nil {
			if left == 0 {
				ans = append(ans, append([]int{}, path...))
			}
			// 叶子节点无论满足条件与否，都直接退出，防止叶子节点的左右nil节点无用递归
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}

	dfs(root, target)
	return ans
}













