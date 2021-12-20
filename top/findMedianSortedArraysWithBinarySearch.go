package top

// 4. 寻找两个正序数组的中位数
// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/xun-zhao-liang-ge-you-xu-shu-zu-de-zhong-wei-s-114/
func findMedianSortedArraysWithBinarySearch(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	}
	midIndex1, midIndex2 := totalLength/2-1, totalLength/2
	f1 := float64(getKthElement(nums1, nums2, midIndex1+1))
	f2 := float64(getKthElement(nums1, nums2, midIndex2+1))
	return (f1 + f2) / 2.0
}

// 在两个数组中找到第n个大的值
func getKthElement(a []int, b []int, k int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	// a b index
	aIndex, bIndex := 0, 0
	for {
		if aIndex == len(a) {
			return b[bIndex+k-1]
		}
		if bIndex == len(b) {
			return a[aIndex+k-1]
		}
		if k == 1 {
			return min(a[aIndex], b[bIndex])
		}
		half := k / 2
		// 防止越界
		newAIndex := min(aIndex+half, len(a)) - 1
		newBIndex := min(bIndex+half, len(b)) - 1
		pivotA, pivotB := a[newAIndex], b[newBIndex]
		if pivotA <= pivotB {
			k = k - (newAIndex - aIndex + 1)
			aIndex = newAIndex + 1
		} else {
			k = k - (newBIndex - bIndex + 1)
			bIndex = newBIndex + 1
		}
	}
}
