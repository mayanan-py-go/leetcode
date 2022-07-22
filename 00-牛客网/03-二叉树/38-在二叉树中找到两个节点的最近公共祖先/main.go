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
	fmt.Println(lowestCommonAncestor(root1, 5, 9))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func lowestCommonAncestor(root *TreeNode, p int, q int) int {
	var flag bool
	var dfs func(*TreeNode, *[]int, int)
	dfs = func(node *TreeNode, li *[]int, target int) {
		if node == nil {
			return
		}
		*li = append(*li, node.Val)
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

	// 寻找p的路径
	pPath := make([]int, 0)
	dfs(root, &pPath, p)

	// 寻找q的路径
	flag = false
	qPath := make([]int, 0)
	dfs(root, &qPath, q)

	// 寻找最后p和q路径中最后相等的就是最近的祖先
	res := 0
	for i := 0; i < len(pPath) && i < len(qPath); i++ {
		x, y := pPath[i], qPath[i]
		if x == y {
			res = x
		}else{
			break
		}
	}
	return res

}














