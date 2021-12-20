package sort

// 88. 合并两个有序数组
// https://leetcode-cn.com/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums1)-1-i; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j+1], nums1[j] = nums1[j], nums1[j+1]
			}
		}
	}
}

// 这种方式基本属于归并排序的范畴了
func mergeV2(nums1 []int, m int, nums2 []int, n int) {
	dump := make([]int, m+n)
	ans := 0
	for a, b := 0, 0; a < m || b < n; ans++ {
		if a == m {
			dump[ans] = nums2[b]
			b++
		} else if b == n {
			dump[ans] = nums1[a]
			a++
		} else {
			if nums1[a] > nums2[b] {
				dump[ans] = nums2[b]
				b++
			} else {
				dump[ans] = nums1[a]
				a++
			}
		}
	}
	for i, v := range dump {
		nums1[i] = v
	}
}

// 这种方式基本属于归并排序的范畴了
func mergeV3(nums1 []int, m int, nums2 []int, n int) {
	for a, b, ans := m-1, n-1, m+n-1; a >= 0 || b >= 0; ans-- {
		if a < 0 {
			nums1[ans] = nums2[b]
			b--
		} else if b < 0 {
			nums1[ans] = nums1[a]
			a--
		} else {
			if nums1[a] >= nums2[b] {
				nums1[ans] = nums1[a]
				a--
			} else {
				nums1[ans] = nums2[b]
				b--
			}
		}
	}
}
