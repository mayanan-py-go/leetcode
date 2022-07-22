package main

import "fmt"

/*
已知两颗二叉树，将它们合并成一颗二叉树。合并规则是：都存在的结点，就将结点值加起来，否则空的位置就由另一个树的结点来代替。例如：
两颗二叉树是:
 */

func main() {
	root1 := &TreeNode{Val: 1}
	n1 := &TreeNode{Val: 2}
	n2 := &TreeNode{Val: 2}
	root1.Left, root1.Right = n1, n2

	root2 := &TreeNode{Val: 1}
	n3 := &TreeNode{Val: 5}
	root2.Left, root2.Right = n3, nil

	fmt.Println(mergeTrees(root1, root2))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	head := &TreeNode{Val: t1.Val+t2.Val}
	head.Left = mergeTrees(t1.Left, t2.Left)
	head.Right = mergeTrees(t1.Right, t2.Right)
	return head
}












