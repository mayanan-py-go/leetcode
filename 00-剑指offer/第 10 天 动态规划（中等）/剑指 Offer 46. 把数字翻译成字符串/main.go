package main

import (
	"fmt"
	"strconv"
)

/*
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。
一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
示例：
输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
*/

func main() {
	fmt.Println(translateNum(12258))
}

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	src := strconv.Itoa(num)
	var ret = 1
	p, q := 0, 1
	for i := 2; i <= len(src); i++ {
		p, q = q, ret
		prev := src[i-2 : i]
		if prev >= "10" && prev <= "25" {
			ret += p
		}
	}
	return ret
}
