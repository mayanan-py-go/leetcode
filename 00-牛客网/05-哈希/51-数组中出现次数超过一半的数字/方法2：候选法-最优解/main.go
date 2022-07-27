package main

import "fmt"

/*
给一个长度为 n 的数组，数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
例如输入一个长度为9的数组[1,2,3,2,2,2,5,4,2]。由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2。
数据范围：n \le 50000n≤50000，数组中元素的值 0 \le val \le 100000≤val≤10000
要求：空间复杂度：O(1)，时间复杂度 O(n)O(n)
*/

func main() {
	fmt.Println(MoreThanHalfNum_Solution([]int{2,2,4,2,5,2,6,2,7}))
}
func MoreThanHalfNum_Solution(numbers []int) int {
	// cond: 候选人，cnt：投票次数
	cond, cnt := -1, 0
	for _, v := range numbers {
		if cnt == 0 {
			cond = v
			cnt ++
		}else if v == cond{
			cnt++
		}else {
			cnt--
		}
	}
	return cond
}










