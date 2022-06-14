package main

import "fmt"

func main() {
	nums := []int{-4,-1,1,2}
	ret := threeSumClosest(nums, 1)
	fmt.Println(ret)
}

/*
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
*/

// 先排序 nums = [-4,-1,1,2]

func threeSumClosest(nums []int, target int) int {
	result := nums[0] + nums[1] + nums[2]

	// 排序nums
	quickSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums)-2; i++{
		l, r := i+1, len(nums)-1
		for l < r{
			sum := nums[i] + nums[l] + nums[r]
			if sum > target{  // r往左移动
				r--
			}else if sum < target{  // l往右移动
				l++
			}else{
				return sum  // 如果相等了就是最接近的值
			}
			// result保存和target相差最小的值
			if distance(sum, target) < distance(result, target){
				result = sum
			}
		}
	}

	return result
}

func distance(a, b int) int{
	if a > b{
		return a - b
	}
	return b - a
}

func quickSort(nums []int, l, r int) {
	if l < r{
		mid := partition(nums, l, r)
		quickSort(nums, l, mid-1)
		quickSort(nums, mid+1, r)
	}
}

func partition(nums []int, l, r int) (mid int){
	tmp := nums[l]
	for l < r{
		for l < r && nums[r] >= tmp{  // 从右边找比tmp小的
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] <= tmp{  // 从左边找比tmp大的
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = tmp  // 将tmp归位
	return l  // 返回l和r重合的位置,也就是mid
}







