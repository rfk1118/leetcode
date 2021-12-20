package sort

import "math/rand"

// 169. 多数元素
// https://leetcode-cn.com/problems/majority-element/
// 使用map的时间复杂度为n，但是空间复杂度为n
func majorityElement(nums []int) int {
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v] = m[v] + 1
	}
	maxCount, result := 0, 0
	for k, v := range m {
		if v > maxCount {
			maxCount = v
			result = k
		}
	}
	return result
}

// 超出时间限制，哈哈哈哈
// 如果空间复杂度为 O(1)，则必须让数组有序，这样才行
func majorityElementV2(nums []int) int {
	// 有序
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums[len(nums)/2]
}

// 随机化，居然还有这种骚套路
func majorityElementV3(nums []int) int {
	var count func([]int, int) int
	count = func(numbers []int, i2 int) int {
		ans := 0
		for _, v := range numbers {
			if v == i2 {
				ans++
			}
		}
		return ans
	}
	majorityCount := len(nums) / 2
	for true {
		i := rand.Intn(len(nums))
		// int candidate = nums[randRange(rand, 0, nums.length)];
		candidate := nums[i]
		if count(nums, candidate) > majorityCount {
			return candidate
		}
	}
	return -1
}

func majorityElementV4(nums []int) int {
	var spiltAndMerge func(nums []int, low, high int) int
	var countData func(nums []int, data int, low, high int) int
	countData = func(nums []int, data, low, high int) int {
		ans := 0
		for i := low; i < high; i++ {
			if nums[i] == data {
				ans++
			}

		}
		return ans
	}
	spiltAndMerge = func(nums []int, low, high int) int {
		if low == high {
			return nums[high]
		}
		mid := (high-low)/2 + low
		left := spiltAndMerge(nums, low, mid)
		right := spiltAndMerge(nums, mid+1, high)
		leftC := countData(nums, left, low, high)
		rightC := countData(nums, right, low, high)
		if leftC > rightC {
			return left
		}
		return right

	}
	return spiltAndMerge(nums, 0, len(nums)-1)
}

func majorityElementV5(nums []int) int {
	ans, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			ans = v
		}
		if ans == v {
			count++
		} else {
			count--
		}
	}
	return ans
}
