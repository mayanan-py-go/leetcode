package main

import (
	"fmt"
)

/*

*/

func main() {
	fmt.Println(fib(10))
}
// 1 1 2 3 5 8 13 21
func fib(n int) int {
	if n < 2 {
		return n
	}
	const mod = 1e9+7
	first, second := 0, 1
	for i := 2; i <= n; i++ {
		second = first + second
		first = (second - first) % mod
	}
	return second
}













