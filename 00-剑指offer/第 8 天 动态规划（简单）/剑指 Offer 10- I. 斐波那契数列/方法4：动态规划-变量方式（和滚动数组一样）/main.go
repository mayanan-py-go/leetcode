package main

import "fmt"

/*
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
*/

func main() {
	fmt.Println(fib(5))
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	const mod = 1e9 + 7
	f1, f2 := 0, 1
	var sum int
	for i := 2; i <= n; i++ {
		sum = (f1 + f2) % mod
		f1, f2 = f2, sum
	}
	return sum
}
