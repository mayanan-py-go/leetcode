package main

/*
一个整型数组里除了两个数字只出现一次，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字。
数据范围：数组长度 2\le n \le 10002≤n≤1000，数组中每个数的大小 0 < val \le 10000000<val≤1000000
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
提示：输出时按非降序排列。
*/

func main() {

}
func FindNumsAppearOnce(array []int) []int {
	m1 := make(map[int]int)
	for _, v := range array {
		if _, ok := m1[v]; ok {
			m1[v] += 1
		}else{
			m1[v] = 1
		}
	}

	ret := make([]int, 2)
	index := 0
	for k, v := range m1 {
		if v == 1 {
			ret[index] = k
			index++
		}
	}
	if ret[0] > ret[1] {
		ret[0], ret[1] = ret[1], ret[0]
	}
	return ret
}











