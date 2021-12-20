package leetcode

// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
// 145. 二叉树的后序遍历
func postorderTraversal(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	var res []int
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if nil == node {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return res
}
