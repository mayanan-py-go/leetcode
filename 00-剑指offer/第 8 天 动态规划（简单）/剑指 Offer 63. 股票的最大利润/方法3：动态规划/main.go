package main

import "fmt"

/*
假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
核心：如果我们是在历史最低点买进的，那么我今天卖出能赚多少钱，当我们考虑完所有天数之时，我们就得到了最好的答案
https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/
*/

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
func maxProfit(prices []int) int {
	var minPrice int = 1e9
	var maxProfit int
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
