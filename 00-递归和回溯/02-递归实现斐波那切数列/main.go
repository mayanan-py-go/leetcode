package main

import "fmt"

/*

*/

func main() {
	fmt.Println(fib(3))
}
// 1 1 2 3 5 8 13 21
func fib(n int) int {
	/*
	时间复杂度：O(2^n)
	空间复杂度：O(n)
	 */
	if n <= 2 {
		return 1
	}
	n1 := fib(n-1)
	n2 := fib(n-2)
	return n1 + n2
}

// 递归调用的空间复杂度 = 递归深度 * 每次调用所需的辅助空间
// 出现了特别多的重复计算

// 递归是自顶向下的调用过程，动态规划是自底向上的计算过程






