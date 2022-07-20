package main

/*
在一个二维数组array中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
[
[1,2,8,9],
[2,4,9,12],
[4,7,10,13],
[6,8,11,15]
]
给定 target = 7，返回 true。
给定 target = 3，返回 false。
数据范围：矩阵的长宽满足 0 \le n,m \le 5000≤n,m≤500 ， 矩阵中的值满足 0 \le val \le 10^90≤val≤10
进阶：空间复杂度 O(1)O(1) ，时间复杂度 O(n+m)O(n+m)
 */

func main() {

}
func Find(target int, array [][]int) bool {
	n := len(array)
	if n < 0 {
		return false
	}
	columns := len(array[0])-1
	m := 0
	for {
		// 处理边界条件
		if m >= n || columns < 0 {
			break
		}
		if array[m][columns] < target {
			m++
		}else if array[m][columns] > target {
			columns--
		}else{
			return true
		}
	}
	return false
}












