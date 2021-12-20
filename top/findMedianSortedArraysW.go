package top

import "math"

// left_A            |          right_A
// A[0], A[1], ..., A[i-1]  |  A[i], A[i+1], ..., A[m-1]

// 4. 寻找两个正序数组的中位数
// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/xun-zhao-liang-ge-you-xu-shu-zu-de-zhong-wei-s-114/
func findMedianSortedArraysW(nums1 []int, nums2 []int) float64 {
	// 让小数组在前面
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	// max
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// min
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	// length
	m, n := len(nums1), len(nums2)
	// 小数组
	left, right := 0, m
	//  median1：前一部分的最大值
	//  median2：后一部分的最小值
	median1, median2 := 0, 0

	for left <= right {
		// 前一部分包含 nums1[0 .. i-1] 和 nums2[0 .. j-1]
		// 后一部分包含 nums1[i .. m-1] 和 nums2[j .. n-1]
		// 求mid，left right 都会动
		i := (left + right) / 2
		j := (m+n+1)/2 - i
		nums_im1 := math.MinInt32
		if i != 0 {
			nums_im1 = nums1[i-1]
		}
		nums_i := math.MaxInt32
		if i != m {
			nums_i = nums1[i]
		}
		nums_jm1 := math.MinInt32
		if j != 0 {
			nums_jm1 = nums2[j-1]
		}
		nums_j := math.MaxInt32
		if j != n {
			nums_j = nums2[j]
		}
		if nums_im1 <= nums_j {
			median1 = max(nums_im1, nums_jm1)
			median2 = min(nums_i, nums_j)
			left = i + 1
		} else {
			right = i - 1
		}
	}
	if (m+n)%2 == 0 {
		return float64(median1+median2) / 2.0
	}
	return float64(median1)
}
