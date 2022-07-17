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
	空间复杂度：O(n)
	 */
	if n <= 1 {
		return 1
	}
	return n + sum(n-1)
}

/*
注意：如果递归调用没有终止，将会一直消耗栈空间
最终导致栈内存溢出，stack overflow
所以必须要有一个明确的结束递归的条件, 也叫边界条件，递归基
 */

/*
使用递归不是为了求得最优解，是为了简化解决问题思路，代码会更简洁
递归求出来的很有可能不是最优解，也有可能是最优解
 */














