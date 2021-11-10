package leetcode

// https://leetcode-cn.com/problems/same-tree/
// 100. 相同的树
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) && p.Val == q.Val
}
