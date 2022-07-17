package main

import (
	"fmt"
)

/*

*/

func main() {
	fmt.Println(fib(35))
}
// 1 1 2 3 5 8 13 21
func fib(n int) int {
	if n < 2 {
		return n
	}
	const mod = 1e9+7
	var arr = make([]int, n+1)
	arr[0], arr[1] = 0, 1
	for i := 2; i <= n; i++ {
		arr[i] = (arr[i-1] + arr[i-2]) % mod
	}
	return arr[n]
}












