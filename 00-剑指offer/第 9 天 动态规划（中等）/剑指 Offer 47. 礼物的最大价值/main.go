package main

import "fmt"

/*
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
*/

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(maxValue(grid))
}
func maxValue(grid [][]int) int {
	row, column := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += max(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[row-1][column-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
