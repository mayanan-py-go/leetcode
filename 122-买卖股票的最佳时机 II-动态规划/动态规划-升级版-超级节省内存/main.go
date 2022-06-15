package main

import "fmt"

/*
给你一个整数数组 prices ，其中prices[i] 表示某支股票第 i 天的价格。
在每一天，你可以决定是否购买和/或出售股票。你在任何时候最多只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
返回 你能获得的 最大 利润。
leetcode链接地址：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
*/

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
