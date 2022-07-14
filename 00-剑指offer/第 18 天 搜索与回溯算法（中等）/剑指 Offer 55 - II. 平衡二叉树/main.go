package main

/*
输入一个二叉树的根节点，判断该二叉树是不是平衡二叉树，如果某二叉树中的任意节点的左右子树深度相差不超过1，那么它就是一颗平衡二叉树
https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/
*/

func main() {

}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left) - height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
func height(root *TreeNode) int {
	 if root == nil {
		 return 0
	 }
	 return max(height(root.Left), height(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(x int) int {
	if x < 0 {
		return -1*x
	}
	return x
}











