package leetcode

// https://leetcode-cn.com/problems/leaf-similar-trees/
// 872. 叶子相似的树
// 又被指针坑了
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var temp01 []int
	var temp02 []int
	var dfs func(node *TreeNode, t *[]int)
	dfs = func(node *TreeNode, t *[]int) {
		if nil == node {
			return
		}
		if node.Left == nil && node.Right == nil {
			*t = append(*t, node.Val)
			return
		}
		dfs(node.Left, t)
		dfs(node.Right, t)
	}
	dfs(root1, &temp01)
	dfs(root2, &temp02)
	if len(temp01) != len(temp02) {
		return false
	}
	for index := range temp01 {
		if temp01[index] != temp02[index] {
			return false
		}
	}
	return true
}
