package main

/*
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。
如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。
https://leetcode.cn/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
*/

func main() {

}
func verifyPostorder(postorder []int) bool {
	return recur(postorder, 0, len(postorder)-1)
}
func recur(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}
	p := i
	for postorder[p] < postorder[j] {
		p++
	}
	m := p  // m左边是左子树，右边包含m是右子树
	for postorder[p] > postorder[j] {
		p++
	}
	return p == j && recur(postorder, i, m-1) && recur(postorder, m, j-1)
}









