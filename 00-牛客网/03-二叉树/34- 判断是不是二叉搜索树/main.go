package main

/*
给定一个二叉树根节点，请你判断这棵树是不是二叉搜索树。
二叉搜索树满足每个节点的左子树上的所有节点均严格小于当前节点且右子树上的所有节点均严格大于当前节点。
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
func isValidBST(root *TreeNode) bool {
	// 中序遍历
	pre := -2<<32
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !dfs(node.Left) {
			return false
		}
		if node.Val < pre {
			return false
		}
		pre = node.Val
		return dfs(node.Right)
	}
	return dfs(root)
}













