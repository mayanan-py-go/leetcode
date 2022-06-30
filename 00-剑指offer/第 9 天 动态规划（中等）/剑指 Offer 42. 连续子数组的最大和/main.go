package main

import "fmt"

/*
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
要求时间复杂度为O(n)。
https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
*/

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
func maxSubArray(nums []int) int {
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	return maxNum
}
