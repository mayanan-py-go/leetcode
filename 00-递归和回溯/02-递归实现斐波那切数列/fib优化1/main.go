package main

import (
	"fmt"
	"time"
)

/*

*/

func main() {
	start := time.Now()
	fmt.Println(fib(35))
	fmt.Println(time.Now().Sub(start))
}
// 1 1 2 3 5 8 13 21
func fib(n int) int {
	if n < 2 {
		return n
	}
	var arr = make([]int, n+1)
	arr[0], arr[1] = 0, 1
	return fib1(n, arr)
}
const mod = 1e9+7
func fib1(n int, arr []int) int {
	if arr[n] == 0 && n != 0 {
		arr[n] = (fib1(n-1, arr) + fib1(n-2, arr)) % mod
	}
	return arr[n]
}












