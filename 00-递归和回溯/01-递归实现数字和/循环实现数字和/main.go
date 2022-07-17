package main

import "fmt"

/*

*/

func main() {
	fmt.Println(sum(4))
}
func sum(n int) int {
	/*
	时间复杂度：O(n)
	空间复杂度：O(1)
	也算是对递归的一个优化
	 */
	var ret int
	for i := 1; i <= n; i++ {
		ret += i
	}
	return ret
}












