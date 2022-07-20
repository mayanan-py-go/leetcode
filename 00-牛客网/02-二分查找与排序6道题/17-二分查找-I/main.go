package main

import "fmt"

/*
请实现无重复数字的升序数组的二分查找
给定一个 元素升序的、无重复数字的整型数组 nums 和一个目标值 target ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标（下标从 0 开始），否则返回 -1
数据范围：0 \le len(nums) \le 2\times10^50≤len(nums)≤2×10
  ， 数组中任意值满足 |val| \le 10^9∣val∣≤10
进阶：时间复杂度 O(\log n)O(logn) ，空间复杂度 O(1)O(1)
 */

func main() {
	fmt.Println(search([]int{}, 0))
}

func search(nums []int, target int) int {
	// 二分查找
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left+right)/2
		if nums[mid] == target {
			return mid
		}else if nums[mid] < target {
			left = mid+1
		}else{
			right = mid-1
		}
	}
	return -1
}














