package main

import "fmt"

/*
其实分两种情况讨论就好了：
当n == 1时，直接将盘子从A移动到C就好了
当n > 1时，可以拆分成3大步骤
将n-1个盘子移动到B, 在将n移动到C, 最后将n-1移动到C
 */

func main() {
	fmt.Println(hanota([]int{2, 1, 0}, []int{}, []int{}))
}
func hanota(A []int, B []int, C []int) []int {
	n := len(A)
	move(n, &A, &B, &C)
	return C
}
func move(n int, A, B, C *[]int) {
	if n == 1 {
		*C = append(*C, (*A)[len(*A)-1])
		*A = (*A)[:len(*A)-1]
	}else {
		move(n-1, A, C, B)
		*C = append(*C, (*A)[len(*A)-1])
		*A = (*A)[:len(*A)-1]
		move(n-1, B, A, C)
	}
}

















