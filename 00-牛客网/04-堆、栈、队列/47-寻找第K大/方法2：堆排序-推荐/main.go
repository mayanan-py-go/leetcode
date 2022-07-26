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
	// 构建堆
	tmp := (n-2) / 2  // 获取最后一个非叶子节点
	for i := tmp; i >= 0; i-- {
		down(a, i, n-1)
	}

	// 获取第k大的数
	for i := n-1; i >= n-K; i-- {
		a[0], a[i] = a[i], a[0]
		down(a, 0, i-1)
	}
	return a[n-K]
}

func down(li []int, low, high int) {
	tmp := li[low]
	left := 2*low+1
	for left <= high {
		// 有右孩子且比左孩子大
		if left+1 <= high && li[left+1] > li[left] {
			left++
		}
		if li[left] > tmp {
			li[low] = li[left]
			low = left
			left = 2*low+1
		}else{
			break
		}
	}
	li[low] = tmp
}














