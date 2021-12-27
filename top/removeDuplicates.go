package top

// 26. 删除有序数组中的重复项
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
// 给你一个有序数组 nums
// 请你原地删除重复出现的元素
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	current, count := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[current] != nums[i] {
			current++
			count++
			nums[current] = nums[i]
		}
	}
	return count
}

//
func removeDuplicatesV2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	slow, fast := 1, 1
	for ; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
