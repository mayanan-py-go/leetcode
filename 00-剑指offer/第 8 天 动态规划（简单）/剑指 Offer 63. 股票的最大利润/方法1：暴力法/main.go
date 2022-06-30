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
	// 核心点：后面的数要大于前面的数，因为股票是按照时间排序的
	var ans int
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			ans = max(ans, prices[j]-prices[i])
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
