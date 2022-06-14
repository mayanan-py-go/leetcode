package main

import (
	"fmt"
)

func main() {
	s := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(s))
}

func maxArea(height []int) int {
	start, area, end := 0, 0, len(height) - 1
	for start < end{
		minHeaght := func(x, y int) int {
			if x > y{
				return y
			}
			return x
		}(height[start], height[end])
		tmp := (end - start) * minHeaght
		if area < tmp{
			area = tmp
		}
		// 从两侧往中间移动，那头小哪头往里移动
		if height[start] > height[end]{
			end--
		}else{
			start++
		}
	}
	return area
}

/*
输入：[1,8,6,2,5,4,8,3,7]
输出：49
*/

// 方法1，功能能实现，但是时间复杂度是O(n^2)导致了leetcode超出时间限制
/*
func maxArea(height []int) int {
	length := len(height)
	n, l := 0, 0
	for i := 0; i < length-1; i++{
		for ii := i+1; ii < length; ii++{
			// 计算高度最小值
			min := func(x, y int) int {
				if x > y{
					return y
				}
				return x
			}(height[i], height[ii])
			l = min * (ii - i)  // 计算面积
			if l > n{n=l}
		}
	}

	return n
}
*/