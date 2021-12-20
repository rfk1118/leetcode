package sort

// 242. 有效的字母异位词
// https://leetcode-cn.com/problems/valid-anagram/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int, 0)
	for _, v := range s {
		m[v] = m[v] + 1
	}
	for _, v := range t {
		m[v] = m[v] - 1
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
