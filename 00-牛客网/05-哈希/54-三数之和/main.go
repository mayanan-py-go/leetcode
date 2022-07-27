package main

import "sort"

func main() {

}
func threeSum(num []int) [][]int {
	sort.Ints(num)

	ret := make([][]int, 0)
	for i := 0; i < len(num); i++ {
		if num[i] > 0 {
			return ret
		}
		// 跳过大循环
		if i > 0 && num[i] == num[i-1] {
			continue
		}
		l, r := i+1, len(num)-1
		for l < r {
			sum := num[i] + num[l] + num[r]
			if sum == 0 {
				ret = append(ret, []int{num[i], num[l], num[r]})
				// 跳过小循环
				for l < r && num[l] == num[l+1] {
					l++
				}
				for l < r && num[r] == num[r-1] {
					r--
				}
				l++
				r--
			}else if sum < 0 {
				l++
			}else if sum > 0{
				r--
			}
		}
	}
	return ret
}














