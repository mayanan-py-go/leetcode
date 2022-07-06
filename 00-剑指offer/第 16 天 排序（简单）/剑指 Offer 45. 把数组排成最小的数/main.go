package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/

说明：
输出结果可能非常大，所以你需要返回一个字符串而不是整数
拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0
 */


func main() {
	nums := []int{3, 30, 34, 9, 5}
	fmt.Println(minNumber(nums))
}

func minNumber(nums []int) string {
	/*
	时间复杂度：O(nLogN)
	空间复杂度：O(n)
	 */

	var strS = make([]string, len(nums))
	for i, v := range nums {
		strS[i] = strconv.Itoa(v)
	}

	var quickSort func([]string, int, int)
	quickSort = func(li []string, l, r int) {
		if l >= r {
			return
		}
		i, j := l, r
		for i < j {
			for i < j && li[j]+li[i] >= li[i]+li[j] {
				j--
			}
			for i < j && li[i]+li[j] <= li[j]+li[i] {
				i++
			}
			li[i], li[j] = li[j], li[i]
		}

		quickSort(li, l, i-1)
		quickSort(li, i+1, r)
	}

	quickSort(strS, 0, len(strS)-1)
	return strings.Join(strS, "")

}














