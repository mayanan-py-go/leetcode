package main

import "fmt"

/*
阶乘
*/

func main() {
	fmt.Println(factorial(10))
	fmt.Println(add(10))
	fmt.Println(factorial2(10))
}
func factorial2(n int) int {
	// 尾递归法
	var dfs func(int, int) int
	dfs = func(n, result int) int {
		if n == 1 {
			return result
		}
		return dfs(n-1, result*n)
	}
	return dfs(n, 1)
}
func factorial(n int) int {
	// 递归法
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
func add(n int) int {
	// 循环法
	var ret = 1
	for i := 1; i <= n; i++ {
		ret *= i
	}
	return ret
}













