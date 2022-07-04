package main

import "fmt"

func main() {
	a := []int{1, 8, 2, 11, 99, 5, 6, 3, 4, 41, 58}
	//a := []int{1, 8, 2}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func quickSort(li []int, l, r int) {
	// 递归退出条件
	if l >= r {
		return
	}

	tmp := li[l]
	low, high := l, r
	for l < r {
		for l < r && li[r] >= li[l] {  // 从右边找比tmp值小的
			r--
		}
		li[l] = li[r]
		for l < r && li[l] <= li[r] {  // 从左边找比tmp值大的
			l++
		}
		li[r] = li[l]
	}
	// 将tmp值归位
	li[l] = tmp

	quickSort(li, low, l-1)
	quickSort(li, l+1, high)
}













