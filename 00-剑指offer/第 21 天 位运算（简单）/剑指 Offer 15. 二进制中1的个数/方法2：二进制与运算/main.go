package main

import "fmt"

/*
编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为 汉明重量).）。
https://leetcode.cn/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
*/

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}
func hammingWeight(num uint32) int {
	var res uint32
	for num != 0 {
		res += 1&num
		num >>= 1
	}
	return int(res)
}













