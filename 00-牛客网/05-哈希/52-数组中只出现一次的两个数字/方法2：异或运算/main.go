package main

import "fmt"

/*
一个整型数组里除了两个数字只出现一次，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字。
数据范围：数组长度 2\le n \le 10002≤n≤1000，数组中每个数的大小 0 < val \le 10000000<val≤1000000
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
提示：输出时按非降序排列。
*/

func main() {
	fmt.Println(1 ^ 1 ^ 2 ^ 3)
}
func FindNumsAppearOnce(array []int) []int {
	temp := 0
	for _, v := range array {
		temp ^= v
	}
	// 此时temp保存的是两个不相同的唯一数字的异或值
	k := 1
	for k & temp == 0 {
		k <<= 1  // 比较更高位，找到不同值
	}
	res1, res2 := 0, 0
	for _, v := range array {
		if k & v == 0 {
			res1 ^= v
		}else{
			res2 ^= v
		}
	}
	if res1 > res2 {
		return []int{res2, res1}
	}
	return []int{res1, res2}
}











