package main

import "fmt"

func main() {
	nums := []int{-2,-1,-1,1,1,2,2}
	ret := fourSum(nums, 0)
	fmt.Println(ret)
}

/*
输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
*/

func fourSum(nums []int, target int) [][]int {
	// 排序
	quickSort(nums, 0, len(nums)-1)

	// 排序后：[-2 -1 0 0 1 2]
	result := make([][]int, 0)
	for i := 0; i < len(nums)-3 && (nums[i]+nums[i+1]+nums[i+2]+nums[i+3]) <= target; i++{

		// 跳过重复值, 当前值加上三个最大值还小于目标就直接遍历下一次
		if i > 0 && nums[i] == nums[i-1] || (nums[i]+nums[len(nums)-3]+nums[len(nums)-2]+nums[len(nums)-1]) < target {
			continue
		}
		for j := i+1; j < len(nums)-2; j++{
			// 跳过重复值, 当前i值+当前j值加两个最大值如果还是小于目标，直接遍历下一次
			if j > i+1 && nums[j] == nums[j-1] || (nums[i]+nums[j]+nums[len(nums)-2]+nums[len(nums)-1]) < target{
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right{
				sums := nums[i] + nums[j] + nums[left] + nums[right]
				if sums == target{
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					// 跳过重复值
					for left < right && nums[left] == nums[left+1]{
						left++
					}
					for left < right && nums[right] == nums[right-1]{
						right--
					}
					left++
					right--
				}else if sums < target{
					left++
				}else if sums > target{
					right--
				}
			}
		}
	}

	return result
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
