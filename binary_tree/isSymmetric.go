package leetcode

// https://leetcode-cn.com/problems/symmetric-tree/
// 101. 对称二叉树
func isSymmetric(root *TreeNode) bool {
	if nil == root {
		return true
	}
	var isSymmetricHelper func(left, right *TreeNode) bool
	isSymmetricHelper = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if (left == nil && right != nil) || (left != nil && right == nil) {
			return false
		}
		return left.Val == right.Val && isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
	}
	return isSymmetricHelper(root.Left, root.Right)
}
