package sort

import "math"

// 628. 三个数的最大乘积
// https://leetcode-cn.com/problems/maximum-product-of-three-numbers/
func maximumProduct(nums []int) int {
	// 如果数组中全是非负数，则排序后最大的三个数相乘即为最大乘积；
	// 如果全是非正数，则最大的三个数相乘同样也为最大乘积
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	i := len(nums)
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	return max(nums[0]*nums[1]*nums[i-1], nums[i-1]*nums[i-2]*nums[i-3])
}

func maximumProductV2(nums []int) int {
	minI, minII := math.MaxInt64, math.MaxInt64
	maxI, maxII, maxIII := math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v < minI {
			minI, minII = v, minI
			// >=
		} else if v >= minI && v < minII {
			minII = v
		}
		if v > maxI {
			maxI, maxII, maxIII = v, maxI, maxII
		} else if v <= maxI && v > maxII {
			maxII, maxIII = v, maxII
		} else if v <= maxII && v > maxIII {
			maxIII = v
		}
	}
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	return max(minI*minII*maxI, maxI*maxII*maxIII)
}
