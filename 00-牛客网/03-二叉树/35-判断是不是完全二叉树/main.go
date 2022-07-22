package main

import "container/list"

/*
给定一个二叉树，确定他是否是一个完全二叉树。
完全二叉树的定义：若二叉树的深度为 h，除第 h 层外，其它各层的结点数都达到最大个数，
第 h 层所有的叶子结点都连续集中在最左边，这就是完全二叉树。（第 h 层可能包含 [1~2h] 个节点）
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
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := list.List{}
	queue.PushBack(root)
	notComplete := false
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		if node == nil {
			notComplete = true
			continue
		}
		if notComplete {
			return false
		}
		queue.PushBack(node.Left)
		queue.PushBack(node.Right)
	}
	return true
}













