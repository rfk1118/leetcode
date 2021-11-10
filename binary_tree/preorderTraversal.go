package leetcode

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
// 144. 二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	var res []int
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if nil == node {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return res
}
