package main

import "fmt"

/*

*/

func main() {
	fmt.Println(sum(4))
}
func sum(n int) int {
	/*
	时间复杂度：O(1)
	空间复杂度：O(1)
	根据规律优化
	 */
	return (1 + n)*n >> 1
}












