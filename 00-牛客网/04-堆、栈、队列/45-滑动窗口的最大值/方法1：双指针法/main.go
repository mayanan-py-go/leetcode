package main

import "fmt"

/*
给定一个长度为 n 的数组 nums 和滑动窗口的大小 size ，找出所有滑动窗口里数值的最大值。
例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}；
针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个： {[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}，
{2,3,[4,2,6],2,5,1}， {2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。
数据范围： 1 \le size \le n \le 100001≤size≤n≤10000，数组中每个元素的值满足 |val| \le 10000∣val∣≤10000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
*/

func main() {
	fmt.Println(maxInWindows([]int{2, 3, 4, 2, 6, 2, 5, 1}, 3))
}
func maxInWindows(num []int, size int) []int {
	if len(num) == 1 {
		return num
	}
	ret := make([]int, 0)
	for left, right := 0, size-1 ;right < len(num); left, right = left+1, right+1 {
		m := max(num[left:right+1])
		ret = append(ret, m)
	}
	return ret
}
func max(li []int) int {
	m := int(-(1e9+7))
	for _, v := range li {
		if v > m {
			m = v
		}
	}
	return m
}
















