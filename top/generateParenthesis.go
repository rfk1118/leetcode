package top

// 22. 括号生成
// https://leetcode-cn.com/problems/generate-parentheses/
func generateParenthesis(n int) (res []string) {
	var dfs func(left, right int, item string)
	dfs = func(left, right int, item string) {
		if left == right && left == n {
			res = append(res, item)
			return
		}
		if right > left || left > n || right > n {
			return
		}
		dfs(left+1, right, item+"(")
		dfs(left, right+1, item+")")
	}
	dfs(0, 0, "")
	return
}

func generateParenthesisV2(n int) (res []string) {
	item := make([]byte, 0)

	var dfs func(left, right int)
	dfs = func(left, right int) {
		if left == right && left == n {
			temp := make([]byte, len(item))
			copy(temp, item)
			res = append(res, string(item))
			return
		}
		if right > left || left > n || right > n {
			return
		}
		item = append(item, '(')
		dfs(left+1, right)
		item = item[:len(item)-1]

		item = append(item, ')')
		dfs(left, right+1)
		item = item[:len(item)-1]
	}

	dfs(0, 0)
	return
}
