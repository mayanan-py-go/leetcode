package main

import "fmt"

/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
算法的时间复杂度应该为 O(log (m+n)) 。
https://leetcode.cn/problems/median-of-two-sorted-arrays/
*/

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	midIndex := totalLength / 2
	if totalLength%2 == 1 { // 奇数
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else { // 偶数
		return float64(getKthElement(nums1, nums2, midIndex)+getKthElement(nums1, nums2, midIndex+1)) / 2
	}
}

func getKthElement(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		// 如果数组1为空了，那么从数组二中取第k个值
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		// 如果数组2为空了，那么从数组一中取第k个值
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 { // 如果k等于1的时候，直接从两个数组中有效区间内取最小值就是对应的中位数
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
