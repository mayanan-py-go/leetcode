package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

/*
请实现两个函数，分别用来序列化和反序列化二叉树，不对序列化之后的字符串进行约束，
但要求能够根据序列化之后的字符串重新构造出一棵与原二叉树相同的树。
二叉树的序列化(Serialize)是指：把一棵二叉树按照某种遍历方式的结果以某种格式保存为字符串，
从而使得内存中建立起来的二叉树可以持久保存。序列化可以基于先序、中序、后序、层序的二叉树等遍历方式来进行修改，
序列化的结果是一个字符串，序列化时通过 某种符号表示空节点（#）
二叉树的反序列化(Deserialize)是指：根据某种遍历顺序得到的序列化字符串结果str，重构二叉树。
*/

func main() {
	root1 := &TreeNode{Val: 8}
	n1 := &TreeNode{Val: 6}
	n2 := &TreeNode{Val: 9}
	n3 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 7}
	root1.Left, root1.Right = n1, n2
	n1.Left, n1.Right = n3, n4
	fmt.Println(Serialize(root1))

	fmt.Println(Deserialize("8#6#9#5#7#-110#-110#-110#-110#-110#-110#"))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func Serialize(root *TreeNode) string {
	var ret string
	if root == nil {
		return ret
	}
	emptyNode := &TreeNode{Val: -110}
	queue := list.List{}
	queue.PushBack(root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		ret  = ret + strconv.Itoa(node.Val) + "#"
		if node != emptyNode {
			if node.Left != nil {
				queue.PushBack(node.Left)
			}else{
				queue.PushBack(emptyNode)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}else{
				queue.PushBack(emptyNode)
			}
		}
	}
	return ret
}
func Deserialize(s string) *TreeNode {
	if s == "" {
		return nil
	}
	ret := strings.Split(s, "#")
	value0, _ := strconv.Atoi(ret[0])
	root := &TreeNode{Val: value0}
	queue := list.List{}
	queue.PushBack(root)
	for i := 1; i < len(ret)-1; i = i+2 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		// 每次从中取出左右节点对应的value
		a, _ := strconv.Atoi(ret[i])
		b, _ := strconv.Atoi(ret[i+1])
		if a != -110 {
			node.Left = &TreeNode{Val: a}
			queue.PushBack(node.Left)
		}
		if b != -110 {
			node.Right = &TreeNode{Val: b}
			queue.PushBack(node.Right)
		}
	}
	return root
}












