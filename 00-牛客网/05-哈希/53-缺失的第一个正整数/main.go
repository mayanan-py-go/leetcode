package main

import "fmt"

/*
给定一个未排序的整数数组nums，请你找出其中没有出现的最小的正整数
进阶： 空间复杂度 O(1)，时间复杂度 O(n)
*/

func main() {
	fmt.Println(minNumberDisappeared([]int{3, 2, 1}))
}
func minNumberDisappeared(nums []int) int {
	m1 := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m1[nums[i]] = true
	}

	i := 1
	for ;i <= len(nums); i++ {
		if m1[i] {
			continue
		}else{
			return i
		}
	}
	return i
}















