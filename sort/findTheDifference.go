package sort

// 389. 找不同
// https://leetcode-cn.com/problems/find-the-difference/
func findTheDifference(s string, t string) byte {
	m := make(map[rune]int, 0)
	for _, v := range s {
		m[v] = m[v] + 1
	}
	for _, v := range t {
		// 这里有两种情况
		// 1. t中这个字母在s中不存在
		// 2. t中的字母比s中多出现了1次
		if m[v] == 0 {
			return byte(v)
		} else {
			m[v] = m[v] - 1
		}
	}
	return byte(' ')
}

func findTheDifferenceSum(s, t string) byte {
	var sum int
	for _, v := range s {
		sum = sum - int(v)
	}
	for _, v := range t {
		sum = sum + int(v)
	}
	return byte(sum)
}

func findTheDifferenceXor(s, t string) byte {
	var sum int
	for _, v := range s {
		sum = sum ^ int(v)
	}
	for _, v := range t {
		sum = sum ^ int(v)
	}
	return byte(sum)
}

// 这个秀，少了1次循环
func findTheDifferenceXorXiu(s, t string) (diff byte) {
	for i := range s {
		diff ^= s[i] ^ t[i]
	}
	return diff ^ t[len(t)-1]
}
