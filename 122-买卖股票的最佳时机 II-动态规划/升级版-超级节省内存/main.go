package main

import "fmt"

func main() {
	prices := []int{8, 7, 1, 3, 15, 4, 2, 6}
	ret := maxProfit(prices)
	fmt.Println(ret)
}

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, 1)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		// 第i天手里没有股票：昨天没有股票今天也不买股票；昨天有股票今天出售股票
		dp[0][0] = max(dp[0][0], dp[0][1]+prices[i])
		// 第i天手里持有股票：昨天有股票今天不买股票；昨天没有股票今天买股票
		dp[0][1] = max(dp[0][1], dp[0][0]-prices[i])
	}
	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
