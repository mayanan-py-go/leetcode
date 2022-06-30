package main

import "fmt"

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n级的台阶总共有多少种跳法。
答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
*/

func main() {
	fmt.Println(numWays(3))
}
func numWays(n int) int {
	if n < 2 {
		return 1
	}
	const mod = 1e9 + 7
	f1, f2 := 1, 1
	var ret int
	for i := 2; i <= n; i++ {
		ret = (f1 + f2) % mod
		f1 = f2
		f2 = ret
	}
	return ret
}
