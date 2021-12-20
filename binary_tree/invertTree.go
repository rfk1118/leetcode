package leetcode

// https://leetcode-cn.com/problems/invert-binary-tree/
// 226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if nil == root {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
