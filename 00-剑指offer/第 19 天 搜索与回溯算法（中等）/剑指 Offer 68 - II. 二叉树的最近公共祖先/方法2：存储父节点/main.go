package main

/*
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
例如，给定如下二叉搜索树:root =6,2,8,0,4,7,9,null,null,3,5]
*/

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parent := make(map[int]*TreeNode)
	visited := make(map[int]bool)

	// 保存所有节点对应的父节点
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parent[node.Left.Val] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			parent[node.Right.Val] = node
			dfs(node.Right)
		}
	}
	dfs(root)

	// 将p节点添加到访问过的节点中去
	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}

	// 从访问过的节点中查找公共节点
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}










