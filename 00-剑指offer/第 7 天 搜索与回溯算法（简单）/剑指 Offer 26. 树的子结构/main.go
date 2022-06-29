package main

import "fmt"

/*
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
B是A的子结构， 即 A中有出现和B相同的结构和节点值。
https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/
*/

func main() {
	A := &TreeNode{Val: 3}
	n1 := &TreeNode{Val: 4}
	n2 := &TreeNode{Val: 5}
	n3 := &TreeNode{Val: 1}
	n4 := &TreeNode{Val: 2}
	A.Left, A.Right = n1, n2
	n1.Left, n1.Right = n3, n4

	B := &TreeNode{Val: 4, Left: &TreeNode{Val: 1}}

	fmt.Println(isSubStructure(A, B))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func recur(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}
