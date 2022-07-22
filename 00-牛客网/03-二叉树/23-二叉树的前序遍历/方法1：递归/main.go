package main

/*

*/

func main() {

}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func preorderTraversal(root *TreeNode) []int {
	var ret = make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		ret = append(ret, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ret
}














