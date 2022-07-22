package main

/*
操作给定的二叉树，将其变换为源二叉树的镜像。
数据范围：二叉树的节点数 0 \le n \le 10000≤n≤1000 ， 二叉树每个节点的值 0\le val \le 10000≤val≤1000
要求： 空间复杂度 O(n)O(n) 。本题也有原地操作，即空间复杂度 O(1)O(1) 的解法，时间复杂度 O(n)O(n)
 */

func main() {
	root1 := &TreeNode{Val: 1}
	n1 := &TreeNode{Val: 2}
	n2 := &TreeNode{Val: 2}
	root1.Left, root1.Right = n1, n2

	root2 := &TreeNode{Val: 1}
	n3 := &TreeNode{Val: 5}
	root2.Left, root2.Right = n3, nil


}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func Mirror( pRoot *TreeNode ) *TreeNode {
	if pRoot == nil {
		return nil
	}
	left := Mirror(pRoot.Left)
	right := Mirror(pRoot.Right)
	pRoot.Left = right
	pRoot.Right = left
	return pRoot
}













