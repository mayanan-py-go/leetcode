package main

import "fmt"

/*

*/

func main() {
	array := []int{3, 2, 5, 66, 77, 2000, 1, 11, 22, 33, 1, 2, 3, 5}
	insertSort(array)
	fmt.Println(array)
}
func insertSort(li []int) {
	// 插入排序
	length := len(li)
	for x := 1; x < length; x++ {
		tmp := li[x]
		y := x-1
		for y >= 0 && li[y] > tmp {
			li[y+1] = li[y]
			y--
		}
		li[y+1] = tmp
	}
}











