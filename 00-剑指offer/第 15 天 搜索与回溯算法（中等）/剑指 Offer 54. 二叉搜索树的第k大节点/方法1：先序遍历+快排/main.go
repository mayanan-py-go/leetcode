package main

import "fmt"

/*
给定一棵二叉搜索树，请找出其中第 k 大的节点的值。
https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/

示例 1:
 */


func main() {
	root := &TreeNode{Val: 3}
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 4}
	n3 := &TreeNode{Val: 2}

	root.Left, root.Right = n1, n2
	n1.Right = n3

	fmt.Println(kthLargest(root, 1))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func kthLargest(root *TreeNode, k int) int {
	/*
	时间复杂度：O(nLogN)
	空间复杂度：O(n)
	 */
	var ret = make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		ret = append(ret, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	quickSort(ret, 0, len(ret)-1)

	return ret[k-1]
}
func quickSort(li []int, l, r int) {
	if l >= r {
		return
	}
	low, high := l, r
	tmp := li[l]
	for l < r {
		for l < r && li[r] <= tmp {  // 从右边找比tmp大的数
			r--
		}
		li[l] = li[r]
		for l < r && li[l] >= tmp {  // 从左边找比tmp小的数
			l++
		}
		li[r] = li[l]
	}
	li[l] = tmp

	quickSort(li, low, l-1)
	quickSort(li, r+1, high)

}













