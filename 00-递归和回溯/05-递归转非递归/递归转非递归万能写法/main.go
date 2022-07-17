package main

import (
	"container/list"
	"fmt"
)

/*
递归转非递归的万能方法：
自己维护一个栈，来保存参数和局部变量
但是空间复杂度依然没有得到优化
 */

func main() {
	log(4)
}
type Frame struct {
	N int
	V int
}
func log(n int) {
	// 递归转非递归（for循环）写法
	var arr = list.List{}
	for ; n>0; n--{
		arr.PushBack(Frame{N: n, V: 10+n})
	}
	for arr.Len() > 0 {
		frame := arr.Remove(arr.Back()).(Frame)
		fmt.Println(frame.V)
	}
}











