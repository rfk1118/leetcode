package top

// 28. 实现 strStr()
// https://leetcode-cn.com/problems/implement-strstr/
func strStr(haystack string, needle string) int {
	var handler func(h string, n string, currentIndex int) int
	handler = func(h string, n string, currentIndex int) int {
		if len(h) == 0 && len(n) == 0 {
			return 0
		}
		if len(n) > len(h) {
			return -1
		}
		allMatch := true
		for j := 0; j < len(n); j++ {
			allMatch = h[j] == n[j] && allMatch
		}
		if allMatch {
			return currentIndex
		}
		return handler(h[1:], n, currentIndex+1)
	}
	return handler(haystack, needle, 0)
}
