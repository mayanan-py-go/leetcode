package main

/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。
https://leetcode.cn/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
*/

func main() {
	myPow(2, 8)
}
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1/quickMul(x, -n)
}
func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y*y
	}
	return y*y*x
}











