package leetcode

// https://leetcode-cn.com/problems/increasing-order-search-tree/submissions/
// 897. 递增顺序搜索树
// 另外一个使用
func increasingBST(root *TreeNode) *TreeNode {
	resTemp := &TreeNode{}
	current := resTemp
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if nil == node {
			return
		}
		dfs(node.Left)
		current.Right = node
		current = node
		node.Left = nil
		dfs(node.Right)
	}
	dfs(root)
	return resTemp.Right
}
