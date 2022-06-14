package main

import "fmt"

func main() {
	nums := []int{-1,0,1,2,-1,-4}
	ret := threeSum(nums)
	fmt.Println(ret)
}

func threeSum(nums []int) (results [][]int) {
	// 冒泡排序
	//bubbleSort(nums)  // 执行效率表较慢

	// 快排
	quickSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums); i++{
		// 1. 如果最左边的值大于0，那么三个数的和加起来一定大于0，那么直接返回
		if nums[i] > 0{
			return
		}
		// 2. 跳过重复值（大循环）
		// 如果当前值和上一个相等，则跳过进行下一次循环
		if i > 0 && nums[i] == nums[i-1]{
			continue  // 相同的值得出的结果列表相同就没必要遍历了
		}
		// 3. 获取左指针和右指针
		l, r := i+1, len(nums)-1
		// 4. 遍历计算结果为0的三元组
		for l < r{
			sums := nums[i] + nums[l] + nums[r]
			if sums == 0{
				results = append(results, []int{nums[i], nums[l], nums[r]})
				// 跳过重复值（小循环）
				for l < r && nums[l] == nums[l+1]{
					l++
				}
				for l < r && nums[r] == nums[r-1]{
					r--
				}
				l++
				r--
			}else if sums < 0{
				l++
			}else if sums > 0{
				r--
			}
		}
	}

	return
}

func bubbleSort(nums []int) {
	// 冒泡排序
	for i := 0; i < len(nums) - 1; i++{
		for j := i + 1; j < len(nums); j++{
			if nums[i] > nums[j]{
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func quickSort(nums []int, l, r int) {
	// 快速排序
	if l < r{
		mid := partition(nums, l, r)
		quickSort(nums, l, mid-1)
		quickSort(nums, mid+1, r)
	}
}

func partition(nums []int, l, r int) (mid int) {
	// 一片一片的进行排序
	tmp := nums[l]
	for l < r{
		for l < r && nums[r] >= tmp{  // 从右面找比tmp小的数
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] <= tmp{  // 从左边找比tmp大的数
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = tmp  // 将tmp归位
	return l  // 返回l或r都可以，因为最终l和r碰上了，位置是一致的
}









