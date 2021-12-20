package sort

// 645. 错误的集合
// https://leetcode-cn.com/problems/set-mismatch/
func findErrorNums(nums []int) []int {
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v] = m[v] + 1
	}

	var ans []int = make([]int, 2)
	for i := 1; i <= len(nums); i++ {
		if m[i] == 2 {
			ans[0] = i
		}
		if m[i] == 0 {
			ans[1] = i
		}
	}
	return ans
}

func findErrorNumsXor(nums []int) []int {
	xor := 0
	for _, v := range nums {
		xor = xor ^ v
	}
	n := len(nums)
	for i := 1; i <= n; i++ {
		xor ^= i
	}

	//  a ^ b
	// 得到 a b
	lowbit := xor & -xor
	num1, num2 := 0, 0
	for _, v := range nums {
		if v&lowbit == 0 {
			num1 ^= v
		} else {
			num2 ^= v
		}
	}
	for i := 1; i <= n; i++ {
		if i&lowbit == 0 {
			num1 ^= i
		} else {
			num2 ^= i
		}
	}
	// 这里是来判断那个是重复的值的
	for _, v := range nums {
		if v == num1 {
			return []int{num1, num2}
		}
	}
	return []int{num2, num1}
}
