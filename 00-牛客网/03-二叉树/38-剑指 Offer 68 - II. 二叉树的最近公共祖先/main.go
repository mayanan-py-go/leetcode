package main

import "fmt"

/*
给定一棵二叉树(保证非空)以及这棵树上的两个节点对应的val值 o1 和 o2，请找到 o1 和 o2 的最近公共祖先节点。
数据范围：树上节点数满足 1 \le n \le 10^5 \1≤n≤10
   , 节点值val满足区间 [0,n)
要求：时间复杂度 O(n)O(n)
注：本题保证二叉树中每个节点的val值均不相同。
 */

func main() {
	root1 := &TreeNode{Val: 8}
	n1 := &TreeNode{Val: 6}
	n2 := &TreeNode{Val: 9}
	n3 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 7}
	root1.Left, root1.Right = n1, n2
	n1.Left, n1.Right = n3, n4
	fmt.Println(lowestCommonAncestor(root1, n4, n3))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	flag := false
	var dfs func(*TreeNode, *[]*TreeNode, int)
	dfs = func(node *TreeNode, li *[]*TreeNode, target int) {
		if node == nil {
			return
		}
		*li = append(*li, node)
		if node.Val == target {
			flag = true
			return
		}
		dfs(node.Left, li, target)
		dfs(node.Right, li, target)
		if flag {
			return
		}
		// 回溯
		*li = (*li)[:len(*li)-1]
	}

	// p的路径
	pPath := make([]*TreeNode, 0)
	dfs(root, &pPath, p.Val)

	flag = false
	// q的路径
	qPath := make([]*TreeNode, 0)
	dfs(root, &qPath, q.Val)

	// 从p和q的路径中寻找它们的最近公共祖先
	var res *TreeNode
	for i := 0; i < len(pPath) && i < len(qPath); i++ {
		x, y := pPath[i], qPath[i]
		if x.Val == y.Val {
			res = x
		} else {
			break
		}
	}
	return res

}













