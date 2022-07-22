package main

import "fmt"

/*
给定一个二叉树root和一个值 sum ，判断是否有从根节点到叶子节点的节点值之和等于 sum 的路径。
1.该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
2.叶子节点是指没有子节点的节点
3.路径只能从父节点到子节点，不能从子节点到父节点
4.总节点数目为n

 */

func main() {
	root := &TreeNode{Val: 1}
	n1 := &TreeNode{Val: 2}
	n2 := &TreeNode{Val: 3}
	n3 := &TreeNode{Val: 4}
	n4 := &TreeNode{Val: 5}
	root.Left, root.Right = n1, n2
	n1.Left, n1.Right = n3, n4
	fmt.Println(hasPathSum(root, 8))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && sum-root.Val == 0 {
		return true
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}














