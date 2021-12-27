package top

// 14. 最长公共前缀
// https://leetcode-cn.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	lcp := func(a, b string) string {
		minl := min(len(a), len(b))
		index := 0
		for ; index < minl && a[index] == b[index]; index++ {
			// 这里啥都不写感觉怪怪的
		}
		return a[:index]
	}
	prefix := strs[0]
	length := len(strs)
	for i := 0; i < length; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func longestCommonPrefixV2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func longestCommonPrefixV3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	lcp := func(a, b string) string {
		minl := min(len(a), len(b))
		index := 0
		for ; index < minl && a[index] == b[index]; index++ {
			// 这里啥都不写感觉怪怪的
		}
		return a[:index]
	}
	var middleHandler func(strs []string, low, high int) string
	middleHandler = func(strs []string, low, high int) string {
		if low == high {
			return strs[low]
		}
		mid := low + (high-low)/2
		leftStr := middleHandler(strs, low, mid)
		rightStr := middleHandler(strs, mid+1, high)
		return lcp(leftStr, rightStr)
	}
	return middleHandler(strs, 0, len(strs)-1)
}

func longestCommonPrefixV4(strs []string) string {
	if nil == strs || len(strs) == 0 {
		return ""
	}
	// find minLength
	minLength := len(strs[0])
	for _, s := range strs {
		if len(s) < minLength {
			minLength = len(s)
		}
	}
	isCommonPrefix := func(mid int) bool {
		str0, count := strs[0][:mid], len(strs)
		for i := 0; i < count; i++ {
			if strs[i][:mid] != str0 {
				return false
			}
		}
		return true
	}
	low, high := 0, minLength
	for low < high {
		mid := (high-low+1)/2 + low
		if isCommonPrefix(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return strs[0][:low]
}
