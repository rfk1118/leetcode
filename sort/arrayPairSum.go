package sort

// 561. 数组拆分 I
// https://leetcode-cn.com/problems/array-partition-i/
func arrayPairSum(nums []int) int {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	sum := 0
	for i := 0; i < len(nums); i = i + 2 {
		sum += nums[i]
	}
	return sum
}
