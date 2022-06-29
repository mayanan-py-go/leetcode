package main

import "fmt"

/*
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
*/

func main() {
	root := &TreeNode{Val: 3}
	n1 := &TreeNode{Val: 9}
	n2 := &TreeNode{Val: 20}
	n3 := &TreeNode{Val: 15}
	n4 := &TreeNode{Val: 7}
	root.Left, root.Right = n1, n2
	n2.Left, n2.Right = n3, n4

	fmt.Println(levelOrder(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)

	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := make([]*TreeNode, 0)
		for len(q) > 0 {
			curNode := q[0]
			ret[i] = append(ret[i], curNode.Val)
			if curNode.Left != nil {
				p = append(p, curNode.Left)
			}
			if curNode.Right != nil {
				p = append(p, curNode.Right)
			}
			q = q[1:]
		}
		q = p
	}
	return ret
}
