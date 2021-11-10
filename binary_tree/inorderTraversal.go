package leetcode

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
// 94. 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	var r []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if nil == node {
			return
		}
		inorder(node.Left)
		r = append(r, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return r
}
