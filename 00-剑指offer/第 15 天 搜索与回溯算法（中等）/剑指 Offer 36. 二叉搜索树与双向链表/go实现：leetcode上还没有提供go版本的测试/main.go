package main

import "fmt"

/*
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。
https://leetcode.cn/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/
 */


func main() {
	root := &TreeNode{Val: 4}
	n1 := &TreeNode{Val: 2}
	n2 := &TreeNode{Val: 5}
	n3 := &TreeNode{Val: 1}
	n4 := &TreeNode{Val: 3}

	root.Left, root.Right = n1, n2
	n1.Left, n1.Right = n3, n4

	fmt.Println(tree2linkedList(root).Right.Right.Right.Right.Right)

}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func tree2linkedList(root *TreeNode) *TreeNode {
	// 中序有顺序的二叉树转双链表, 不借助任何新的节点
	if root == nil {
		return nil
	}
	var head, prev *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)  // 递归左子树
		if prev == nil {
			head = node  // 第一次访问，声明头结点
		}else{
			prev.Right, node.Left = node, prev  // 因为是双链表，修改双向指针
		}
		prev = node  // 修改上一个节点为当前节点
		dfs(node.Right)
	}
	dfs(root)
	head.Left, prev.Right = prev, head  // 首位相连，形成环形双链表
	return head
}














