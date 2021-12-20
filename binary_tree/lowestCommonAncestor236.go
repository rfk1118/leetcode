package leetcode

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 236. 二叉树的最近公共祖先
func lowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 第一次在父亲集合
	if left != nil && right != nil {
		return root
	}
	// 已经在同一个子树了
	if left != nil {
		return left
	}
	return right
}
