package main

import "fmt"

/*
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
*/

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(findNumberIn2DArray(matrix, 20))
}
func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 从左下角开始查找
	i, j := len(matrix)-1, 0
	for {
		if i < 0 || j >= len(matrix[0]) {
			return false
		}
		flag := matrix[i][j]
		if flag > target {
			i--
		} else if flag < target {
			j++
		} else {
			return true
		}
	}
}
