package main

import "fmt"

/*
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/
*/

func main() {
	nums := []int{0, 1}
	fmt.Println(missingNumber(nums))
}
func missingNumber(nums []int) int {
	length := len(nums)
	// 首个数字必须是0
	if nums[0] != 0 {
		return 0
	}
	// 长度是1，首个数字是0的时候需要单独处理
	if length == 1 && nums[0] == 0 {
		return 1
	}
	// 中间缺了数字的情况
	for i := 0; i < length-1; i++ {
		if nums[i+1]-nums[i] == 2 {
			return nums[i+1] - 1
		}
	}
	// 缺最后一个数字的情况
	return nums[length-1] + 1
}
