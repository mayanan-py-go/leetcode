package main

import "fmt"

/*
给定一棵二叉树，判断其是否是自身的镜像（即：是否对称）
例如：                                 下面这棵二叉树是对称的
 */

func main() {
	root := &TreeNode{Val: 1}
	n1 := &TreeNode{Val: 2}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 4}
	n4 := &TreeNode{Val: 4}
	root.Left, root.Right = n1, n2
	n1.Left, n2.Right = n3, n4
	fmt.Println(isSymmetrical(root))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func isSymmetrical(pRoot *TreeNode) bool {
	var recursion func(*TreeNode, *TreeNode) bool
	recursion = func(root1 *TreeNode, root2 *TreeNode) bool {
		// 递归界
		if root1 == nil && root2 == nil {
			return true
		}
		if root1 == nil || root2 == nil || root1.Val != root2.Val {
			return false
		}
		return recursion(root1.Left, root2.Right) && recursion(root1.Right, root2.Left)
	}
	return recursion(pRoot, pRoot)
}












