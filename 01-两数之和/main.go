package main

import "fmt"

func main() {
	s1 := []int{11, 33, 33, 44, 22}
	target := 66
	s2 := twoSum(s1, target)
	fmt.Println(s2)  // [1 2]
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for index, val := range nums{
		if preIndex, ok := m[target - val]; ok{
			return []int{preIndex, index}
		} else {
			m[val] = index
		}
	}
	return []int{}
}








