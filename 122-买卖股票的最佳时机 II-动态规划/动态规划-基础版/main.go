package main

import "fmt"

/*
给你一个整数数组 prices ，其中prices[i] 表示某支股票第 i 天的价格。
在每一天，你可以决定是否购买和/或出售股票。你在任何时候最多只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
返回 你能获得的 最大 利润。
leetcode链接地址：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
*/

func main() {
	prices := []int{7, 8, 5, 3, 6, 40}
	maxProfitRet := maxProfit(prices)
	fmt.Println(maxProfitRet)
}

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = 0          // 第0天交易完成后，手无股票的最大利润
	dp[0][1] = -prices[0] // 第0天交易完成后，手持股票的最大利润
	// 状态转移方程
	for i := 1; i < n; i++ {
		// dp[i][0]: 上一天没有股票这一天也没买，或者上一天持有股票这一天卖了, 求这两种情况下的最大利润即可
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) // 第i天交易完成后，手无股票的最大利润
		// dp[i][1]: 上一天持有股票这一天也没卖，上一天没有股票这一天买了, 求这两种情况下的最大利润即可
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) // 第i天交易完成后，手持股票的最大利润
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
