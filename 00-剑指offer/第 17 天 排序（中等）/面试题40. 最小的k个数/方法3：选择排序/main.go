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
	fmt.Println(getLeastNumbers([]int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}, 8))
}
func getLeastNumbers(arr []int, k int) []int {
	var selectSort func([]int, int)
	selectSort = func(arr []int, k int) {
		length := len(arr)
		for i := 0; i < k; i++ {
			minLoc := i
			for j := i+1; j < length; j++ {
				if arr[j] < arr[minLoc] {
					minLoc = j
				}
			}
			arr[i], arr[minLoc] = arr[minLoc], arr[i]
		}
	}

	selectSort(arr, k)
	fmt.Println(arr)
	return arr[:k]
}
