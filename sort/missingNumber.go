package sort

// 268. 丢失的数字
// https://leetcode-cn.com/problems/missing-number/
// 将数组排序之后，即可根据数组中每个下标处的元素是否和下标相等，得到丢失的数字。
func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1]-1 {
			return nums[i] + 1
		}
	}

	if nums[len(nums)-1] == len(nums) {
		return 0
	}

	return len(nums)
}

func missingNumberV2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != i {
			return i
		}
	}

	return len(nums)
}

func missingNumberV3(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	for i := 0; i <= len(nums); i++ {
		if !m[i] {
			return i
		}
	}
	return len(nums)
}

func missingNumberV4(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans = ans ^ v
	}

	for i := 0; i <= len(nums); i++ {
		ans = ans ^ i
	}
	return ans
}

func missingNumberV5(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum = sum + v
	}
	i := len(nums)
	total := i * (i + 1) / 2
	return total - sum
}
