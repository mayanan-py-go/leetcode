package main

import "fmt"

/*
输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。
https://leetcode.cn/problems/he-wei-sde-liang-ge-shu-zi-lcof/

示例：
输入：nums = [2,7,11,15], target = 9
输出：[2,7] 或者 [7,2]
 */

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
func twoSum(nums []int, target int) []int {
	/* 对撞双指针
	时间复杂度：O(n)
	空间复杂度：O(1)
	 */
	left, right := 0, len(nums)-1
	for left < right {
		tmp := nums[left] + nums[right]
		if tmp > target{
			right--
		}else if tmp < target {
			left++
		}else{
			return []int{nums[left], nums[right]}
		}
	}
	return []int{}
}
















