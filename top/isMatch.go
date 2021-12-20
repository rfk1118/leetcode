package top

// 10. 正则表达式匹配
// https://leetcode-cn.com/problems/regular-expression-matching/
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	match := func(i, j int) bool {
		if i == 0 {
			return false
		}
		// 匹配任意单个字符
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}
	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			//  匹配零个或多个前面的那一个元素
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if match(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if match(i, j) {
				// 不是*直接处理
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	// 这就是走到最后了
	return f[m][n]
}
