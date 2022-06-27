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
	nums1 := []int{1, 2, 3}
	nums1 = append(nums1, 0, 0, 0)
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	var ret = make([]int, m+n)
	index1, index2 := 0, 0
	for i := 0; i < len(ret); i++ {
		if index1 < m && index2 < n {
			if nums1[index1] < nums2[index2] {
				ret[i] = nums1[index1]
				index1++
			} else {
				ret[i] = nums2[index2]
				index2++
			}
		} else if index1 < m {
			ret[i] = nums1[index1]
			index1++
		} else if index2 < n {
			ret[i] = nums2[index2]
			index2++
		}
	}
	copy(nums1, ret)
}
