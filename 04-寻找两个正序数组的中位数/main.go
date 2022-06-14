package main

import "fmt"

func main() {
	ret := findMedianSortedArrays([]int{11, 15, 22}, []int{3, 8, 13, 222})
	fmt.Println(ret)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	nums := make([]int, len(nums1) + len(nums2))
	// 定义两个变量标识nums1和nums2的索引
	i, j := 0, 0

	// 1. 将两两个切片中的值按照正序顺序写入到新建的nums切片中
	for k := 0; k < len(nums); k++{
		if i < m && j < n{  // 代表两个切片中都有值
			// 判断nums1值的值和nums2中的值，哪个小写入到nums切片中
			if nums1[i] < nums2[j]{
				nums[k] = nums1[i]
				i++
			}else{
				nums[k] = nums2[j]
				j++
			}
		}else if i < m{  // 代表nums1切片中有值
			nums[k] = nums1[i]
			i++
		}else if j < n{  // 代表nums2切片中有值
			nums[k] = nums2[j]
			j++
		}
	}

	// 此时nums已经是一个正序的切片，取中位数
	// 分两种情况，nums的长度是奇数或偶数
	if len(nums) % 2 == 0{  // 偶数：需要计算中间两个数的平均数
		return float64(nums[len(nums) / 2] + nums[len(nums) / 2 - 1]) / 2
	}else{  // 奇数：取中间的数即可
		return float64(nums[len(nums) / 2])
	}
}
