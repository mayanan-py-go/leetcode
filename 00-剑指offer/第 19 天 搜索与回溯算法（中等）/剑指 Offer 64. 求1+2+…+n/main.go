package main

import "fmt"

/*
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
https://leetcode.cn/problems/qiu-12n-lcof/
*/

func main() {
	fmt.Println(sumNums(3))
}
func sumNums(n int) int {
	var ans int
	var sum func(int) bool
	sum = func(n int) bool {
		ans += n
		return n > 1 && sum(n-1)
	}
	sum(n)
	return ans
}













