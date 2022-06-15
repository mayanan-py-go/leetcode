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
	ret := 0
	for i := 1; i < n; i++ {
		curProfit := prices[i] - prices[i-1]
		if curProfit > 0 {
			ret += curProfit
		}
	}
	return ret
}
