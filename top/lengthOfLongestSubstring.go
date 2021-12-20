package top

// 3. 无重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	ans := 0
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		m := make(map[rune]bool, 0)
		currentlength := 0
		for j := i; j < len(r); j++ {
			if !m[r[j]] {
				m[r[j]] = true
				currentlength++
			} else {
				break
			}
		}
		if currentlength > ans {
			ans = currentlength
		}
	}
	return ans
}
