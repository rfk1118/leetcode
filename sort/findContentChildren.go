package sort

// 455. 分发饼干
// https://leetcode-cn.com/problems/assign-cookies/
// 哈哈哈，大家都是练习算法的
func findContentChildren(g []int, s []int) int {
	var sort func(nums []int)
	sort = func(nums []int) {
		for i := 0; i < len(nums); i++ {
			min := i
			for j := i + 1; j < len(nums); j++ {
				if nums[j] < nums[min] {
					min = j
				}
			}
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
	sort(g)
	sort(s)
	ans := 0
	// 给x小孩找饼干
	for x, y := 0, 0; x < len(g) && y < len(s); x++ {
		for y < len(s) && g[x] > s[y] {
			y++
		}
		if y < len(s) {
			ans++
			y++
		}
	}
	return ans
}
