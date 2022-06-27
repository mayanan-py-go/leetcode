package main

import "fmt"

/*
统计一个数字在排序数组中出现的次数。
https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
*/

func main() {
	nums := []int{1, 2, 5, 5, 8, 8, 10} // 8
	fmt.Println(search(nums, 8))
}
func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	// 查找right
	for i <= j {
		m := (i + j) / 2
		if nums[m] <= target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	right := i
	// 如果没有找到，直接返回, 此时的nums[j]是查找到target的最后一个值，也就是8在索引5的位置
	if j > 0 && nums[j] != target {
		return 0
	}
	// 查找left
	i = 0 // 此时j等于right-1
	for i <= j {
		m := (i + j) / 2
		if nums[m] >= target {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	left := j
	return right - left - 1
}
