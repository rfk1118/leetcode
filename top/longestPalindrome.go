package top

// 5. 最长回文子串
// https://leetcode-cn.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	le := len(s)
	// 如果长度等于1的话
	if le < 2 {
		return s
	}
	maxLen := 1
	begin := 0
	// dp[i][j] 表示 s[i..j] 是否是回文串
	var dp [][]bool = make([][]bool, 0)
	// 初始化：所有长度为 1 的子串都是回文串
	for i := 0; i < le; i++ {
		var inner []bool = make([]bool, le)
		inner[i] = true
		dp = append(dp, inner)
	}

	// 递推开始
	// 先枚举子串长度
	// 最小的长度为2，最大长度为所有
	for L := 2; L <= le; L++ {
		// 枚举左边界，左边界的上限设置可以宽松一些
		for i := 0; i < le; i++ {
			// 由 L 和 i 可以确定右边界，即 j - i + 1 = L 得
			j := L + i - 1
			if j >= le {
				break
			}
			// 两个边界不相同，那说明这个子串不是回文
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				// 就两个属性，还想等，直接处理
				if j-i < 3 {
					dp[i][j] = true
				} else {
					// 如果这里相等，也就是孩子相等
					dp[i][j] = dp[i+1][j-1]
				}
			}
			// 如果大于了max,记录max和开始
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func longestPalindromeV2(s string) string {
	// 每次都从第i个位置向两头走
	expandAroundCenter := func(s string, left, right int) (int, int) {
		//  left == right ,则这个就是中心，向两边扩散
		//  如果 left = right-1,则是处理abba问题
		for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
		}
		return left + 1, right - 1
	}
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

// todo 马拉车后面看
func longestPalindromev3(s string) string {
	start, end := 0, -1
	// 进行数据扩容
	t := "#"
	for i := 0; i < len(s); i++ {
		t += string(s[i]) + "#"
	}
	t += "#"

	s = t
	arm_len := []int{}
	right, j := -1, -1
	for i := 0; i < len(s); i++ {
		var cur_arm_len int
		if right >= i {
			i_sym := j*2 - i
			min_arm_len := min(arm_len[i_sym], right-i)
			cur_arm_len = expand(s, i-min_arm_len, i+min_arm_len)
		} else {
			cur_arm_len = expand(s, i, i)
		}
		arm_len = append(arm_len, cur_arm_len)
		if i+cur_arm_len > right {
			j = i
			right = i + cur_arm_len
		}
		if cur_arm_len*2+1 > end-start {
			start = i - cur_arm_len
			end = i + cur_arm_len
		}
	}
	ans := ""
	for i := start; i <= end; i++ {
		if s[i] != '#' {
			ans += string(s[i])
		}
	}
	return ans
}

func expand(s string, left, right int) int {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return (right - left - 2) / 2
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
