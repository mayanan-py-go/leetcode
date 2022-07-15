package main

/*
写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。
https://leetcode.cn/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/
*/

func main() {

}
func add(a, b int) int {
	for b != 0 {  // 当进位为0时跳出
		c := (a&b)<<1  // c = 进位
		a ^= b  // a = 非进位和
		b = c  // b = 进位
	}
	return a
}







