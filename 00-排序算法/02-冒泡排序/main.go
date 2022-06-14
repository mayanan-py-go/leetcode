package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 排序前
	var nums = make([]int, 0, 100)
	rand.Seed(time.Now().UnixNano())  // 随机种子
	for i := 0; i < cap(nums); i++{
		nums = append(nums, rand.Intn(100))
	}
	fmt.Println(nums)

	bubbleSort(nums)  // 冒泡排序

	// 排序后
	fmt.Println(nums)
}

func bubbleSort(nums []int) {
	// 冒泡排序
	for i := 0; i < len(nums)-1; i++{
		for j := i+1; j < len(nums); j++{
			if nums[i] > nums[j]{
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
