package main

import "fmt"

/*
前序遍历：ABCDEFGH
中序遍历：BDCEAFHG
后续遍历：DECBHGFA
层次遍历：ABFCGDEH
*/

func main() {
	var nodeA = &binaryNode{Val: 'A'}
	var nodeB = &binaryNode{Val: 'B'}
	var nodeC = &binaryNode{Val: 'C'}
	var nodeD = &binaryNode{Val: 'D'}
	var nodeE = &binaryNode{Val: 'E'}
	var nodeF = &binaryNode{Val: 'F'}
	var nodeG = &binaryNode{Val: 'G'}
	var nodeH = &binaryNode{Val: 'H'}
	nodeA.lChild = nodeB
	nodeA.rChild = nodeF
	nodeB.rChild = nodeC
	nodeC.lChild = nodeD
	nodeC.rChild = nodeE
	nodeF.rChild = nodeG
	nodeG.lChild = nodeH

	// 前序遍历
	//preOrder(nodeA)

	// 中序遍历
	//midOrder(nodeA)

	// 后续遍历
	//backOrder(nodeA)

	// 层次遍历
	layerOrder(nodeA)
}

type binaryNode struct {
	Val    byte
	lChild *binaryNode
	rChild *binaryNode
}

// 前序遍历
func preOrder(node *binaryNode) {
	if node == nil {
		return
	}
	fmt.Println(string(node.Val))
	preOrder(node.lChild)
	preOrder(node.rChild)
}

// 中序遍历
func midOrder(node *binaryNode) {
	if node == nil {
		return
	}
	midOrder(node.lChild)
	fmt.Println(string(node.Val))
	midOrder(node.rChild)
}

// 后续遍历
func backOrder(node *binaryNode) {
	if node == nil {
		return
	}
	backOrder(node.lChild)
	backOrder(node.rChild)
	fmt.Println(string(node.Val))
}

// 层次遍历
func layerOrder(node *binaryNode) {
	if node == nil {
		return
	}
	// 将根节点加入到队列
	l1 := make([]*binaryNode, 0)
	l1 = append(l1, node)
	for len(l1) > 0 {
		head := l1[0]
		fmt.Println(string(head.Val))
		if head.lChild != nil {
			l1 = append(l1, head.lChild)
		}
		if head.rChild != nil {
			l1 = append(l1, head.rChild)
		}
		l1 = l1[1:]
	}
}
