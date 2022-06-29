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
	// 方法3：辅助队列实现
	var s1 = []*TreeNode{root}
	for len(s1) > 0 {
		curNode := s1[0]
		if curNode.Left != nil {
			s1 = append(s1, curNode.Left)
		}
		if curNode.Right != nil {
			s1 = append(s1, curNode.Right)
		}
		curNode.Left, curNode.Right = curNode.Right, curNode.Left
		s1 = s1[1:]
	}
	return root
}
