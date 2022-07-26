package main

import "fmt"

/*
有一个整数数组，请你根据快速排序的思路，找出数组中第 k 大的数。
给定一个整数数组 a ,同时给定它的大小n和要找的 k ，请返回第 k 大的数(包括重复的元素，不用去重)，保证答案存在。
要求：时间复杂度 O(nlogn)O(nlogn)，空间复杂度 O(1)O(1)
数据范围：0\le n \le 10^50≤n≤10
 ， 1 \le K \le n1≤K≤n，数组中每个元素满足 0 \le val \le 10^90≤val≤10
*/

func main() {
	arr := []int{1, 8, 8, 5, 6}
	fmt.Println(findKth(arr, len(arr), 3))
}
func findKth( a []int ,  n int ,  K int ) int {
	quickSort(a, 0, n-1)
	return a[K-1]
}
func quickSort(li []int, left, right int) {
	if left >= right {
		return
	}
	l, r := left, right
	tmp := li[l]
	for l < r {
		for l < r && li[r] <= tmp {  // 从右边找比tmp大的数
			r--
		}
		li[l] = li[r]
		for l < r && li[l] >= tmp {  // 从左边找比tmp小的数
			l++
		}
		li[r] = li[l]
	}
	li[l] = tmp

	// 递归+分治
	quickSort(li, left, l-1)
	quickSort(li, r+1, right)
}














