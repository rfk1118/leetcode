package top

import "math"

// 11. 盛最多水的容器
// https://leetcode-cn.com/problems/container-with-most-water/
// 要解决这个问题，必须两个线很高，并且够宽
// dp 而且只能用备忘录，因为后面的如果比前面的要高，那就是会扩大值
func maxArea(height []int) int {
	max := math.MinInt64
	minIndex, maxIndex := 0, len(height)-1
	minFunc := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	maxFunc := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	for minIndex < maxIndex {
		max = maxFunc(max, minFunc(height[minIndex], height[maxIndex])*(maxIndex-minIndex))
		if height[minIndex] > height[maxIndex] {
			maxIndex--
		} else {
			minIndex++
		}
	}
	return max
}
