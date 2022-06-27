package main

import "fmt"

/*
统计一个数字在排序数组中出现的次数。
https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
*/

func main() {
	nums := []int{1, 3, 5, 5, 8, 8, 10} // 8
	fmt.Println(search(nums, 8))
}
func search(nums []int, target int) int {
	return binarySearch(nums, target) - binarySearch(nums, target-1)
}
func binarySearch(nums []int, target int) (right int) {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] <= target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}
