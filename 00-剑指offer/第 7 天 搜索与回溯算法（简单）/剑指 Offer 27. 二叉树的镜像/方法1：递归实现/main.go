package main

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
	tmp := root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(tmp)
	return root
}
