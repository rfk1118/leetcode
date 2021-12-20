package top

import "math"

// 7. 整数反转
// https://leetcode-cn.com/problems/reverse-integer/
// 输入：x = 123
// 输出：321
func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x = x / 10
		rev = rev*10 + digit
	}
	return
}
