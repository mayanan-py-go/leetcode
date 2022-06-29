package main

import "fmt"

/*
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
*/

func main() {
	// 声明树
	var root = &TreeNode{Val: 3}
	var n1 = &TreeNode{Val: 9}
	var n2 = &TreeNode{Val: 20}
	var n3 = &TreeNode{Val: 15}
	var n4 = &TreeNode{Val: 7}
	root.Left, root.Right = n1, n2
	n2.Left, n2.Right = n3, n4

	// 层次遍历
	fmt.Println(levelOrder(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var ret = make([]int, 0)
	if root == nil {
		return ret
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		head := queue[0]
		ret = append(ret, head.Val)
		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
		queue = queue[1:]
	}
	return ret
}
