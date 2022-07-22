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
func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		ret = append(ret, node.Val)
	}
	dfs(root)
	return ret
}












