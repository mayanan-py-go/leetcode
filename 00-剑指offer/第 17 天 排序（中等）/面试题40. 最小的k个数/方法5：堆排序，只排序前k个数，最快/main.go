package main

import (
	"fmt"
)

/*
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

示例1：
输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
*/

func main() {
	fmt.Println(getLeastNumbers([]int{2, 3, 1, 5, 9, 4, 1, 3}, 3))
}

func getLeastNumbers(arr []int, k int) []int {
	heapSort(arr, k)
	return arr[len(arr)-k:]
}
func heapSort(li []int, k int) {
	// 构建小根堆, 从最后一个非叶子节点进行构建
	tmp := (len(li) - 1 - 1) / 2  // 获取最后一个非叶子节点
	for i := tmp; i >= 0; i-- {
		sift(li, i, len(li)-1)
	}

	// 堆构建完成，开始进行排序（排序的过程本质上就是构建堆的过程）
	// 注意：因为我们只需要前k小的数，所以我们只需要排完前k个就可以了
	for i := len(li)-1; i >= len(li)-k; i-- {
		li[0], li[i] = li[i], li[0]
		sift(li, 0, i-1)
	}

}
func sift(li []int, low, high int) {
	tmp := li[low]    // 暂存当前节点数据
	left := 2*low + 1 // 左孩子
	for left <= high {  // 只要有左孩子就遍历
		if left+1 <= high && li[left+1] < li[left] { // 如果有右孩子，而且右孩子比左孩子小
			left += 1
		}
		if li[left] < tmp {
			li[low] = li[left]
			low = left // 往下看一层
			left = 2*low + 1  // 同时修改左孩子
		}else {
			break
		}
	}
	li[low] = tmp  // 把tmp归位
}














