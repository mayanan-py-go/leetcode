package main

import "fmt"

/*

*/

func main() {
	arr := []int{1, 8, 6, 3, 5}
	fmt.Println(GetLeastNumbers_Solution(arr, 2))
}
func GetLeastNumbers_Solution(li []int, k int) []int {
	// 从最后一个非叶子节点构建堆
	tmp := (len(li)-2) / 2
	for i := tmp; i >= 0; i-- {
		down(li, i, len(li)-1)
	}

	// 通过构建好的堆进行排序
	for i := len(li)-1; i >= len(li)-k; i-- {
		li[0], li[i] = li[i], li[0]
		down(li, 0, i-1)
	}
	return li[len(li)-k:]
}
func down(li []int, low, high int) {
	// 暂存当前节点的值
	tmp := li[low]
	left := 2*low + 1  // 计算左孩子
	for left <= high {
		// 如果有右孩子而且右孩子比左孩子小,获取右孩子
		if left+1 <= high && li[left+1] < li[left] {
			left++
		}
		if li[left] < tmp {
			li[low] = li[left]
			low = left
			left = 2*low + 1
		}else{
			break
		}
	}
	li[low] = tmp
}














