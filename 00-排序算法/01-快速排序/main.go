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

	quickSort(nums, 0, len(nums) - 1)  // 快排

	// 排序后
	fmt.Println(nums)
}

func quickSort(nums []int, l, r int) {
	if l < r{
		mid := partition(nums , l, r)
		quickSort(nums, l, mid-1)
		quickSort(nums, mid+1, r)
	}
}

func partition(nums []int, l, r int) (mid int){
	tmp := nums[l]
	for l < r{
		for l < r && nums[r] >= tmp{  // 从右边找比tmp小的值放到左边
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] <= tmp{  // 从左边找比tmp大的值放到右边
			l++
		}
		nums[r] = nums[l]
	}
	// 将tmp值归位
	nums[l] = tmp
	return l  // 返回l和r其实是一样的，因为他俩已经在同一个位置了
}

