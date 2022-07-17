package main

import "fmt"

/*

*/

func main() {
	fmt.Println(fib(10))
}
func fib(n int) int {

	var dfs func(int, int, int) int
	dfs = func(n, first, second int) int {
		if n <= 1 {
			return first
		}
		return dfs(n-1, second, first+second)
	}

	return dfs(n, 1, 1)
}













