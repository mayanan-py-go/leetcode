package main

import "fmt"

/*

*/

func main() {
	array := []int{3, 2, 5, 66, 77, 2000, 1, 11, 22, 33, 1, 2, 3, 5}
	fmt.Println(getLeastNumbers(array, 3))
}
func getLeastNumbers(arr []int, k int) []int {
	return heapSort(arr, k)
}
func heapSort(li []int, k int) []int {
	/* 小根堆：降序排序 */
	// 从最后一个非叶子节点构建堆
	tmp := (len(li)-2)/2
	for i := tmp; i >= 0; i-- {
		down(li, i, len(li)-1)
	}

	// 堆构建完成之后进行降序排序（本质上就是构建小根堆的过程）
	for i := len(li)-1; i >= len(li)-k; i-- {
		li[0], li[i] = li[i], li[0]
		down(li, 0, i-1)
	}
	return li[len(li)-k:]
}
func down(li []int, low, high int) {
	tmp := li[low]
	left := 2*low+1
	for left <= high {
		// 有右孩子而且右孩子小于左孩子
		if left+1 <= high && li[left+1] < li[left] {
			left++
		}
		if li[left] < tmp {
			li[low] = li[left]
			low = left  // 往下看一层
			left = 2*low+1
		}else {
			break
		}
	}
	li[low] = tmp
}












