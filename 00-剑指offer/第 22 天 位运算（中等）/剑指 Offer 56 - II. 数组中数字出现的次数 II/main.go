package main

import "fmt"

/*
在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。
https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/
*/

func main() {
	fmt.Println(singleNumber([]int{1, 4, 1, 1, 2, 4, 4}))
}
func singleNumber(nums []int) int {
	var ones, twos int
	for _, v := range nums {
		ones = ones ^ v & ^twos
		twos = twos ^ v & ^ones
	}
	return ones
}

/* 上面代码的解释
0 ^ x = x,
x ^ x = 0；
x & ~x = 0,
x & ~0 =x;
那么就是很好解释下面的代码了。一开始a = 0, b = 0;
x第一次出现后，a = (a ^ x) & ~b的结果为 a = x, b = (b ^ x) & ~a的结果为此时因为a = x了，所以b = 0。
x第二次出现：a = (a ^ x) & ~b, a = (x ^ x) & ~0, a = 0; b = (b ^ x) & ~a 化简， b = (0 ^ x) & ~0 ,b = x;
x第三次出现：a = (a ^ x) & ~b， a = (0 ^ x) & ~x ,a = 0; b = (b ^ x) & ~a 化简， b = (x ^ x) & ~0 , b = 0; 所以出现三次同一个数，a和b最终都变回了0.
只出现一次的数，按照上面x第一次出现的规律可知a = x, b = 0;因此最后返回a.
 */







