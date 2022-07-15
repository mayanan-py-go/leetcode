package main

import (
	"fmt"
	"strconv"
)

/*
编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为 汉明重量).）。
https://leetcode.cn/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
*/

func main() {
	fmt.Println(hammingWeight(10))
}
func hammingWeight(num uint32) int {
	ret := strconv.FormatUint(uint64(num), 2)
	var rt int
	for _, v := range ret {
		if v == '1' {
			rt++
		}
	}
	return rt
}













