package main

import "fmt"

/*
给定节点数为 n 的二叉树的前序遍历和中序遍历结果，请重建出该二叉树并返回它的头结点。
例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建出如下图所示。

重点：前序遍历找到二叉树的根，中序遍历找到左子树和右子树
*/

func main() {
	a1 := []int{1,2,4,7,3,5,6,8}
	a2 := []int{4,7,2,1,5,3,8,6}
	fmt.Println(reConstructBinaryTree(a1, a2))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	n := len(pre)
	m := len(vin)
	if n == 0 || m == 0 {
		return nil
	}
	root := &TreeNode{Val: pre[0]}
	for i := 0; i < m; i++ {
		// 找到中序遍历的前序第一个元素
		if vin[i] == root.Val {
			// 构建左子树
			root.Left = reConstructBinaryTree(pre[1:i+1], vin[:i])
			// 构建右子树
			root.Right = reConstructBinaryTree(pre[i+1:], vin[i+1:])
			break
		}
	}
	return root
}
















