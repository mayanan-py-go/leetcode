package main

import (
	"fmt"
	"math"
)

/*
输入一棵节点数为 n 二叉树，判断该二叉树是否是平衡二叉树。
在这里，我们只需要考虑其平衡性，不需要考虑其是不是排序二叉树
平衡二叉树（Balanced Binary Tree），具有以下性质：它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，
并且左右两个子树都是一棵平衡二叉树。
 */

func main() {
	root1 := &TreeNode{Val: 1}
	n1 := &TreeNode{Val: 2}
	n2 := &TreeNode{Val: 3}
	n3 := &TreeNode{Val: 4}
	n4 := &TreeNode{Val: 5}
	root1.Left, root1.Right = n1, nil
	n1.Left, n1.Right = n2, nil
	n2.Left , n2.Right = n3, nil
	n3.Left, n3.Right = n4, nil
	fmt.Println(IsBalanced_Solution(root1))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func IsBalanced_Solution(pRoot *TreeNode) bool {
	isBalanced := true
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)

		if math.Abs(float64(l-r)) > 1 {
			isBalanced = false
		}

		return max(l, r)+1
	}
	dfs(pRoot)
	return isBalanced
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}














