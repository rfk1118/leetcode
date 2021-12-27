package top

// 17. 电话号码的字母组合
// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	// 初始化数据
	var symbolValues = map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"}}
	length := len(digits)
	// 如果为空字符串，不处理
	if length == 0 {
		return nil
	}
	// 进行分治
	var handler func(low, high int) []string
	handler = func(low, high int) []string {
		if low == high {
			return symbolValues[digits[low]]
		}
		mid := low + (high-low)/2
		// 左边获取到结果
		leftHandler := handler(low, mid)
		// 右边的结果
		rightHandler := handler(mid+1, high)
		// 合并结果
		var ans []string = make([]string, 0)
		for _, l := range leftHandler {
			for _, r := range rightHandler {
				ans = append(ans, l+r)
			}
		}
		return ans
	}
	return handler(0, length-1)
}
