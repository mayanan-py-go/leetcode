package main

import "fmt"

/*
假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/
*/

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
func maxProfit(prices []int) int {
	var minPrice int = 1e9
	var maxProfit = 0
	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		}
		if v-minPrice > maxProfit {
			maxProfit = v - minPrice
		}
	}
	return maxProfit
}
