package leetcode

// https://leetcode-cn.com/problems/unique-binary-search-trees-ii/solution/bu-tong-de-er-cha-sou-suo-shu-ii-by-leetcode-solut/
// 95. 不同的二叉搜索树 II
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	var helper func(s, e int) []*TreeNode
	helper = func(s, e int) []*TreeNode {
		if s > e {
			// nil 和 []*TreeNode{nil} 不一样，要保证能foreach
			return []*TreeNode{nil}
		}
		var allTrees []*TreeNode
		for i := s; i <= e; i++ {
			leftTrees := helper(s, i-1)
			rightTrees := helper(i+1, e)
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					c := &TreeNode{Val: i}
					c.Left = left
					c.Right = right
					allTrees = append(allTrees, c)
				}
			}
		}
		return allTrees
	}
	return helper(1, n)
}
