package top

// 15. 三数之和
// https://leetcode-cn.com/problems/3sum/
func threeSum(nums []int) [][]int {
	var ans [][]int
	count := len(nums)
	// sort
	for i := 0; i < count; i++ {
		min := i
		for j := i + 1; j < count; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}

	// 从第一个元素开始
	for x := 0; x < count; x++ {
		// 主要是与前面元素不能相同
		if x > 0 && nums[x] == nums[x-1] {
			continue
		}
		third := count - 1
		target := -1 * nums[x]
		// y元素从x+1开始
		for y := x + 1; y < count; y++ {
			// 防止便利相同
			if y > x+1 && nums[y] == nums[y-1] {
				continue
			}
			// 从右边开始处理
			for y < third && nums[y]+nums[third] > target {
				third--
			}
			// 没有找到结果
			if y == third {
				break
			}

			if nums[y]+nums[third] == target {
				// 找到结果了存储到数组中
				ans = append(ans, []int{nums[x], nums[y], nums[third]})
			}
		}
	}

	return ans
}
