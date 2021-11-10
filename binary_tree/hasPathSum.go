package leetcode

// https://leetcode-cn.com/problems/path-sum/
// 112. 路径总和
// 根节点到叶子节点
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var hasPathSumHelper func(node *TreeNode, current, targetSum int) bool
	hasPathSumHelper = func(node *TreeNode, current, targetSum int) bool {
		// 这里代表已经走到了叶子节点
		if node.Left == nil && node.Right == nil {
			return current+node.Val == targetSum
		}
		// 下面代表还没走到叶子几点，所以需要向下走
		res := false
		current = current + node.Val
		if node.Left != nil {
			res = res || hasPathSumHelper(node.Left, current, targetSum)
		}
		if node.Right != nil {
			res = res || hasPathSumHelper(node.Right, current, targetSum)
		}
		return res
	}
	return hasPathSumHelper(root, 0, targetSum)
}
