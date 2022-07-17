package main

import "fmt"

/*
在某些时候，也可以重复使用一组相同的变量来保存每个栈帧的内容
 */

func main() {
	log(4)
}

func log(n int) {
	// 这里重复使用变量i保存原来栈帧中的参数
	// 空间复杂度从O(n)变成O(1)
	for i := 1; i <= n; i++ {
		fmt.Println(10+i)
	}
}











