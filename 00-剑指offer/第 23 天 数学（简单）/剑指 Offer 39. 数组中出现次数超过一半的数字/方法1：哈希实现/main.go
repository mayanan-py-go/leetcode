package main

import "fmt"

/*
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
https://leetcode.cn/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
*/

func main() {
	fmt.Println(majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}
func majorityElement(nums []int) int {
	var m1 = make(map[int]int)
	var max, value int
	for _, num := range nums {
		if v, ok := m1[num]; ok {
			m1[num] = v + 1
		}else {
			m1[num] = 1
		}
		if m1[num] > max {
			max, value = m1[num], num
		}
	}
	return value
}








