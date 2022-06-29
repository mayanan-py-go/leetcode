package main

import "fmt"

/*
给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。（深度优先搜索-递归）
https://leetcode.cn/problems/same-tree/
*/

func main() {
	A := &TreeNode{Val: 4}
	n1 := &TreeNode{Val: 1}
	A.Left = n1

	B := &TreeNode{Val: 4, Left: &TreeNode{Val: 2}}

	fmt.Println(isSameTree(A, B))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
