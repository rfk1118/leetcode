package top

// 20. 有效的括号
// https://leetcode-cn.com/problems/valid-parentheses/
func isValid(s string) bool {
	count := len(s)
	if count%2 == 1 {
		return false
	}
	var pairs map[byte]byte = map[byte]byte{
		')': '(', '}': '{', ']': '[',
	}
	stack := []byte{}
	for i := 0; i < count; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
