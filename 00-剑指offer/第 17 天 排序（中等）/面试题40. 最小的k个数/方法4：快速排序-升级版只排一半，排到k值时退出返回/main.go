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
	fmt.Println(getLeastNumbers([]int{2, 3, 1}, 2))
}

func getLeastNumbers(arr []int, k int) []int {
	var quickSort func([]int, int, int, int)
	quickSort = func(li []int, l int, r int, k int) {
		// 切记此递归退出条件
		if l >= r {
			return
		}

		i, j := l, r
		for i < j {
			for i < j && li[j] >= li[i] {
				j--
			}
			for i < j && li[i] <= li[j] {
				i++
			}
			li[i], li[j] = li[j], li[i]
		}

		num := i-l+1
		if k == num {
			return
		}else if k < num {
			quickSort(li, l, i-1, k)
		}else{
			quickSort(li, i+1, r, k-num)
		}
	}

	quickSort(arr, 0, len(arr)-1, k)
	return arr[:k]
}













