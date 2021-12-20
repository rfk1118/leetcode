package sort

// 594. 最长和谐子序列
// https://leetcode-cn.com/problems/longest-harmonious-subsequence/
func findLHS(nums []int) int {
	// insert sort
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; i < len(nums); j++ {
			if nums[j] > nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	ans := 0
	begin := 0
	for end, num := range nums {
		for num-nums[begin] > 1 {
			begin++
		}
		if num-nums[begin] == 1 && end-begin+1 > ans {
			ans = end - begin + 1
		}
	}
	return ans
}

func findLHSMap(nums []int) int {
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v] = m[v] + 1
	}

	max := 0
	for k, v := range m {
		i := m[k+1]
		if i > 0 && v+i > max {
			max = v + i
		}
	}
	return max
}
