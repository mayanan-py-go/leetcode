package main

import "fmt"

/*

*/

func main() {
	array := []int{3, 2, 5, 66, 77, 2000, 1, 11, 22, 33, 1, 2, 3, 5}
	heapSort(array)
	fmt.Println(array)
}
func heapSort(li []int) {
	/* 大根堆：升序排序 */
	// 从最后一个非叶子节点构建堆
	tmp := (len(li)-1-1)/2
	for i := tmp; i >= 0; i-- {
		down(li, i, len(li)-1)
	}

	// 构建好堆之后开始进行排序(排序的过程本质上就是构建堆的过程)
	for i := len(li)-1; i >= 0; i-- {
		li[0], li[i] = li[i], li[0]
		down(li, 0, i-1)
	}
}
func down(li []int, low, high int) {
	tmp := li[low]  // 暂存当前节点
	left := 2*low+1  // 左孩子
	for left <= high {  // 如果有左孩子的话
		if left+1 <= high && li[left+1] > li[left] {  // 如果有右孩子，而且右孩子比较大的话
			left++
		}
		if li[left] > tmp {
			li[low] = li[left]
			low = left  // 往下看一层
			left = 2*low + 1  // 同时修改左孩子
		}else{
			break
		}
	}
	li[low] = tmp
}












