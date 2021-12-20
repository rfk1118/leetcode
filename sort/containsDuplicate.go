package sort

// 217. 存在重复元素
// https://leetcode-cn.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v] = m[v] + 1
	}
	for _, v := range m {
		if v > 1 {
			return true
		}
	}
	return false
}

func containsDuplicateV2(nums []int) bool {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == v {
				return true
			}
		}
	}
	return false
}
