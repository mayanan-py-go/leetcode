package main

import "fmt"

/*
给你两个按 非递减顺序 排列的整数数组nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，
其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
https://leetcode.cn/problems/merge-sorted-array/
*/

func main() {
	nums1 := []int{1, 2, 5, 6}
	nums1 = append(nums1, 0)
	m := 4
	nums2 := []int{3}
	n := 1
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	index1, index2 := m-1, n-1
	var i = len(nums1) - 1
	for {
		if index1 < 0 {
			copy(nums1[:index2+1], nums2[:index2+1])
			break
		}
		if index2 < 0 {
			break
		}
		if nums1[index1] < nums2[index2] {
			nums1[i] = nums2[index2]
			i--
			index2--
		} else {
			nums1[i] = nums1[index1]
			i--
			index1--
		}
	}
}
