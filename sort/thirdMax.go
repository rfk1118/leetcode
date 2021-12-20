package sort

import "math"

// 414. 第三大的数
// https://leetcode-cn.com/problems/third-maximum-number/
func thirdMax(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	pre, ans := math.MinInt64, 0
	for _, v := range nums {
		if v != pre {
			ans++
			if ans == 3 {
				return v
			}
			pre = v
		}
	}
	return nums[0]

}

func thirdMaxV2(nums []int) int {
	max, mid, min := math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v > max {
			max, mid, min = v, max, mid
		} else if v < max && v > mid {
			mid, min = v, mid
		} else if v > min && v < mid {
			min = v
		}
	}
	if min == math.MinInt64 {
		return max
	}
	return max
}
