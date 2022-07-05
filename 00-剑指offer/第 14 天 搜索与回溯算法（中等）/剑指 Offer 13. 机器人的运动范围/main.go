package main

import "fmt"

/*
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，
它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/


 */


func main() {
	fmt.Println(movingCount(3, 2, 17))
}

func movingCount(m int, n int, k int) int {
	// 储存结果
	res := 0

	// 用来做标记的数组
	var arr [][]bool
	for i := 0; i < m; i++ {
		arr = append(arr, []bool{})
		for j := 0; j < n; j++ {
			arr[i] = append(arr[i], false)
		}
	}

	dfs(0,0, m, n, k, arr, &res)
	return res
}
func dfs(i, j, m, n, k int, arr [][]bool, res *int){
	// 基本判断+判断是否走过这条路
	if i < 0 || i >= m || j < 0 || j >= n || arr[i][j] {
		return
	}
	// 没有走过，先标记，在判断是否符合题意
	arr[i][j] = true
	// 判断行列坐标之和是否大于k
	sum := i%10 + j%10 + i/10 + j/10
	if sum > k {
		return
	}
	// 符合题意数量+1
	*res++
	// 直接大范围撒网（优化：搜索方向缩减为向下和向右，不必再向上和向左搜索了）
	dfs(i+1, j, m, n, k, arr, res)
	//dfs(i-1, j, m, n, k, arr, res)
	dfs(i, j+1, m, n, k, arr, res)
	//dfs(i, j-1, m, n, k, arr, res)
}












