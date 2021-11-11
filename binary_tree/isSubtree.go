package leetcode

// 572. 另一棵树的子树
// https://leetcode-cn.com/submissions/detail/237716107/
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	var check func(a, b *TreeNode) bool
	check = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		// 相等的话，则需要都相等
		if a.Val == b.Val {
			return check(a.Left, b.Left) && check(a.Right, b.Right)
		}
		return false
	}
	return check(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}
