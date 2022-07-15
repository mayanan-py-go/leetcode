package main

import "fmt"

/*
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
*/

func main() {
	fmt.Println(singleNumbers([]int{4, 1, 4, 6}))
}
func singleNumbers(nums []int) []int {
	var ret int
	for _, v := range nums {
		ret ^= v  // ret的最终答案就是那两个只出现一次数的异或结果
	}
	// 找到ret二进制中第几位是1
	var target = 1
	for (target&ret) == 0 {
		target <<= 1
	}
	// 分组
	var a, b int
	for _, v := range nums {
		if (v&target) == 0 {  // 第一组
			a ^= v
		}else {  // 第二组
			b ^= v
		}
	}
	return []int{a, b}
}









