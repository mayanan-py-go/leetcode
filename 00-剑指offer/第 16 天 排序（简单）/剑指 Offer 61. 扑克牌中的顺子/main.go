package main

import "fmt"

/*
从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。
2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。
https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof/

示例：
输入: [0,0,1,2,5]
输出: True
*/

func main() {
	nums := []int{2, 0, 5, 0, 1}
	fmt.Println(isStraight(nums))
}

func isStraight(nums []int) bool {
	/* 充分条件
	1. 除大小王外，所有牌无重复
	2. 记录5张牌中的最大值，最小值（大小王除外），需满足max-min<5
	*/

	var res = make(map[int]bool, 0)
	max, min := 0, 14
	for _, v := range nums {
		if v == 0 {
			continue
		}
		if res[v] {
			return false
		}
		res[v] = true
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max-min < 5
}






