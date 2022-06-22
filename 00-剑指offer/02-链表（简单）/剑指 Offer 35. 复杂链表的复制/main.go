package main

import "fmt"

/*
请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点
还有一个 random 指针指向链表中的任意节点或者 null。
https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/
*/

func main() {
	// 回溯 + 哈希表 实现
	n1 := &Node{Val: 12}
	n1.Next = &Node{Val: 13}
	n1.Random = &Node{Val: 14}
	n1.Next.Next = &Node{Val: 15}
	n1.Random = &Node{Val: 16}

	fmt.Printf("%p\n", n1)
	fmt.Printf("%p\n", copyRandomList(n1))
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cacheNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, has := cacheNode[node]; has {
		return n
	}
	newNode := &Node{Val: node.Val}
	cacheNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}
func copyRandomList(head *Node) *Node {
	cacheNode = map[*Node]*Node{}
	return deepCopy(head)
}
