package main

import "container/list"

/*
请完成一个函数，输入一个二叉树，该函数输出它的镜像。
https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/
注意：本题与主站 226 题相同：https://leetcode-cn.com/problems/invert-binary-tree/
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 辅助栈
	var stack1 = new(list.List)
	stack1.PushBack(root)
	for stack1.Len() > 0 {
		node := stack1.Remove(stack1.Back()).(*TreeNode)
		if node.Left != nil {
			stack1.PushBack(node.Left)
		}
		if node.Right != nil {
			stack1.PushBack(node.Right)
		}
		node.Left, node.Right = node.Right, node.Left
	}
	return root
}
